<template>
  <div class="scene-container">
    <div class="sidebar">
      <h2>阵营与队伍</h2>

      <div class="player-switch">
        <span
          v-for="player in players"
          :key="player.id"
          :class="['player-tab', { active: currentPlayerId === player.id }]"
          @click="currentPlayerId = player.id; selectedTeamId = null"
        >
          {{ player.name }}
        </span>
      </div>

      <div class="action-mode">
        <h3>行为模式</h3>
        <div class="mode-switch">
          <span
            :class="['mode-tab', { active: actionMode === 'march' }]"
            @click="actionMode = 'march'"
          >
            行军
          </span>
          <span
            :class="['mode-tab', { active: actionMode === 'paint' }]"
            @click="actionMode = 'paint'"
          >
            染色
          </span>
        </div>
      </div>

      <div class="team-list">
        <h3>队伍（最多 5 支）</h3>
        <div
          v-for="team in currentPlayerTeams"
          :key="team.id"
          :class="['team-item', { selected: selectedTeamId === team.id }]"
          @click="selectTeam(team)"
        >
          <div class="team-header">
            <span class="team-type">
              <span class="team-avatar">{{ getTeamAvatar(team.type) }}</span>
              {{ team.type }}
            </span>
            <span class="team-status">
              {{ team.hexIndex === null ? '未上场' : '已在地图' }}
            </span>
          </div>
          <div class="team-sub">
            {{ teamLabel(team) }}
          </div>
        </div>
      </div>

      <div class="tips">
        <p>操作说明：</p>
        <ul>
          <li>1. 选择行为模式：行军或染色。</li>
          <li>2. 选择一支队伍。</li>
          <li>3. 点击格子：</li>
          <li class="indent">- <strong>行军</strong>：只移动队伍，不改变格子颜色。</li>
          <li class="indent">- <strong>染色</strong>：移动队伍并染色，但只能从已染色的格子染色。</li>
          <li>4. 城池由 7 个黄色格子组成（中心 + 外围一圈 6 个）。</li>
          <li>5. 当队伍移动到城池中心格子时，会占领该城池。</li>
          <li>6. 按住鼠标左键拖动地图。</li>
        </ul>
      </div>
    </div>

    <div
      class="hex-grid"
      ref="mapContainer"
      @mousedown="handleMouseDown"
      @mousemove="handleMouseMove"
      @mouseup="handleMouseUp"
      @mouseleave="handleMouseUp"
      @touchstart="handleTouchStart"
      @touchmove="handleTouchMove"
      @touchend="handleTouchEnd"
      @touchcancel="handleTouchEnd"
    >
      <svg
        ref="mapSvg"
        :width="svgWidth"
        :height="svgHeight"
        :viewBox="viewBox"
        @click="handleGridClick"
      >

        <!-- 六边形地块 -->
        <g v-for="(hex, index) in hexagons" :key="index">
          <polygon
            :points="hex.points"
            :fill="hex.color"
            stroke="#333"
            :stroke-width="5"
            @click="(e) => handleHexClick(e, hex)"
            @touchstart="(e) => handleHexTouchStart(e, hex)"
            @touchend="(e) => handleHexTouchEnd(e, hex)"
            :data-index="index"
          />
          <!-- 染色倒计时显示 -->
          <template v-for="(task, taskKey) in paintingTasks" :key="taskKey">
            <template v-if="task.hexIndex === hex.index">
              <rect
                :x="hex.centerX - hexagonSize * 0.4"
                :y="hex.centerY - hexagonSize * 0.6"
                :width="hexagonSize * 0.8"
                :height="hexagonSize * 0.4"
                fill="#fff"
                fill-opacity="0.9"
                stroke="#000"
                stroke-width="1"
                rx="4"
                pointer-events="none"
              />
              <text
                :x="hex.centerX"
                :y="hex.centerY - hexagonSize * 0.35"
                :font-size="hexagonSize * 0.35"
                text-anchor="middle"
                fill="#ef4444"
                font-weight="bold"
                pointer-events="none"
              >
                {{ getPaintingCountdown(taskKey) }}s
              </text>
            </template>
          </template>
        </g>

        <!-- 城池轮廓（方便视觉识别） -->
        <g v-for="city in cities" :key="city.id">
          <polygon
            :points="cityOutlinePoints(city)"
            fill="none"
            stroke="#000"
            :stroke-width="strokeWidth * 1.5"
          />
        </g>

        <!-- 队伍标记（使用头像） -->
        <g v-for="team in teams" :key="team.id">
          <template v-if="team.hexIndex !== null">
            <circle
              :cx="getTeamDisplayX(team)"
              :cy="getTeamDisplayY(team)"
              :r="avatarSize * 0.55"
              :fill="playerById(team.playerId)?.unitColor"
              stroke="#000"
              :stroke-width="strokeWidth * 0.8"
            />
            <text
              :x="getTeamDisplayX(team)"
              :y="getTeamDisplayY(team) + avatarSize / 2.5"
              :font-size="avatarSize * 0.85"
              text-anchor="middle"
              fill="#000"
              pointer-events="none"
            >
              {{ getTeamAvatar(team.type) }}
            </text>
          </template>
        </g>
      </svg>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue';

// ======= 地图与六边形网格 =======
let gridSize = 5000; // 支持50x50网格（基础值）

// 根据设备类型调整网格大小
const adjustGridSizeForDevice = () => {
  if (isMobileDevice()) {
    gridSize = 15000; // 手机端使用更大的网格来容纳更多格子
  } else {
    gridSize = 5000; // web端保持原有大小
  }
};
const hexagonSize = ref(80); // 单个六边形边长（响应式）

// 检测是否为移动设备
const isMobileDevice = () => {
  return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent) ||
         (window.innerWidth <= 768 && window.innerHeight <= 1024);
};

// 根据设备类型调整格子大小和其他参数
const adjustHexagonSizeForDevice = () => {
  if (isMobileDevice()) {
    hexagonSize.value = 80 * 3; // 手机端格子为web端的三倍
    adjustGridSizeForDevice(); // 调整网格大小
    adjustSvgSizeForDevice(); // 调整SVG尺寸
    adjustViewBoxForDevice(); // 调整viewBox
  } else {
    hexagonSize.value = 80; // web端保持原有大小
    gridSize = 5000; // 重置为默认值
    svgWidth.value = 5000;
    svgHeight.value = 4500;
    viewBox.value = '0 0 5000 4500';
  }
};
const hexagons = ref([]);
const rowsCount = ref(0);
const colsCount = ref(0);

// ======= 地图拖动与缩放 =======
const mapContainer = ref(null);
const mapSvg = ref(null);
let svgWidth = ref(5000);
let svgHeight = ref(4500);

// 根据设备类型调整SVG尺寸
const adjustSvgSizeForDevice = () => {
  if (isMobileDevice()) {
    svgWidth.value = 15000;
    svgHeight.value = 13500;
  } else {
    svgWidth.value = 5000;
    svgHeight.value = 4500;
  }
};
let viewBox = ref('0 0 5000 4500');

// 根据设备类型调整viewBox
const adjustViewBoxForDevice = () => {
  if (isMobileDevice()) {
    viewBox.value = '0 0 15000 13500';
  } else {
    viewBox.value = '0 0 5000 4500';
  }
};
const isDragging = ref(false);
const dragStart = ref({ x: 0, y: 0 });
const dragDistance = ref(0);
const viewBoxOffset = ref({ x: 0, y: 0 });
const strokeWidth = computed(() => 1.5);
const avatarSize = computed(() => hexagonSize.value);

// ======= 队伍头像映射 =======
const teamAvatars = {
  '骑兵': '🐴',
  '盾兵': '🛡️',
  '枪兵': '🔱',
  '弓兵': '🏹',
  '投石车': '⚔️'
};

// ======= WebSocket 连接 =======
// WebSocket连接状态和相关变量

const ws = ref(null);
const wsConnected = ref(false);
const wsReconnectAttempts = ref(0);
const maxReconnectAttempts = 5;
const reconnectInterval = ref(null);

// WebSocket配置
const WS_URL = 'ws://localhost:7604/array/ws'; // 可以根据需要修改URL

// WebSocket消息处理函数
const wsMessageHandlers = {
  // 处理队伍移动消息
  'march': (data) => {
    console.log('收到队伍移动消息:', data);
    // TODO: 处理队伍移动同步
  },

  // 处理染色消息
  'dye': (data) => {
    console.log('收到染色消息:', data);
    // TODO: 处理格子染色同步
  },

  // 处理城池占领消息
  'city_capture': (data) => {
    console.log('收到城池占领消息:', data);
    // TODO: 处理城池占领同步
  },

  'login_success': (data) => {
    console.log('收到登录成功消息:', data);
    // TODO: 处理登录成功同步
  },

  // 处理游戏状态同步
  'welcome': (data) => {
    console.log('收到游戏状态同步:', data);
    // TODO: 处理游戏状态同步
  }
};

// 连接WebSocket
const connectWebSocket = () => {
  if (ws.value && ws.value.readyState === WebSocket.OPEN) {
    return; // 已经连接
  }
  console.log("开始连接");
  try {
    ws.value = new WebSocket(WS_URL);

    ws.value.onopen = () => {
      console.log('WebSocket连接成功');
      wsConnected.value = true;
      wsReconnectAttempts.value = 0;

      // 发送连接成功消息，包含玩家信息
      sendWebSocketMessage('join_game', {
        ClientID: currentPlayerId.value,
        playerName: playerById(currentPlayerId.value)?.name || '未知玩家',
        Data:{},
        Timestamp: Date.now(),
      });
    };

    ws.value.onmessage = (event) => {
      try {
        const message = JSON.parse(event.data);
        console.log('收到WebSocket消息:', message);

        const handler = wsMessageHandlers[message.type];
        if (handler) {
          handler(message.data);
        } else {
          console.warn('未知消息类型:', message.type);
        }
      } catch (error) {
        console.error('解析WebSocket消息失败:', error);
      }
    };

    ws.value.onclose = (event) => {
      console.log('WebSocket连接关闭:', event.code, event.reason);
      wsConnected.value = false;

      // 自动重连逻辑
      if (wsReconnectAttempts.value < maxReconnectAttempts) {
        wsReconnectAttempts.value++;
        console.log(`尝试重连WebSocket (${wsReconnectAttempts.value}/${maxReconnectAttempts})`);

        setTimeout(() => {
          connectWebSocket();
        }, 2000 * wsReconnectAttempts.value); // 递增重连间隔
      }
    };

    ws.value.onerror = (error) => {
      console.error('WebSocket连接错误:', error);
      wsConnected.value = false;
    };

  } catch (error) {
    console.error('创建WebSocket连接失败:', error);
  }
};

// 断开WebSocket连接
const disconnectWebSocket = () => {
  if (reconnectInterval.value) {
    clearInterval(reconnectInterval.value);
    reconnectInterval.value = null;
  }

  if (ws.value) {
    ws.value.close();
    ws.value = null;
  }
  wsConnected.value = false;
  wsReconnectAttempts.value = 0;
};

// 发送WebSocket消息
const sendWebSocketMessage = (type, data) => {
  if (!ws.value || ws.value.readyState !== WebSocket.OPEN) {
    console.warn('WebSocket未连接,无法发送消息');
    return false;
  }

  const message = {
    type,
    data,
    timestamp: Date.now(),
    playerId: currentPlayerId.value
  };

  try {
    ws.value.send(JSON.stringify(message));
    return true;
  } catch (error) {
    console.error('发送WebSocket消息失败:', error);
    return false;
  }
};

// 发送队伍移动消息
const sendTeamMoveMessage = (teamId, fromHexIndex, toHexIndex) => {
  return sendWebSocketMessage('team_move', {
    teamId,
    fromHexIndex,
    toHexIndex,
    playerId: currentPlayerId.value
  });
};

// 发送染色消息
const sendPaintMessage = (hexIndex, playerId, teamId) => {
  return sendWebSocketMessage('paint_hex', {
    hexIndex,
    playerId,
    teamId
  });
};

// 发送城池占领消息
const sendCityCaptureMessage = (cityId, playerId) => {
  return sendWebSocketMessage('city_capture', {
    cityId,
    playerId
  });
};

const getTeamAvatar = (type) => {
  return teamAvatars[type] || '⚔️';
};

// ======= 阵营与队伍数据 =======
const maxTeamsPerPlayer = 5;
const teamTypes = ['骑兵', '盾兵', '枪兵', '弓兵', '投石车'];

const players = ref([
  {
    id: 1,
    name: '阵营 1',
    tileColor: '#4ade80', // 普通地块染色（绿色）
    unitColor: '#16a34a'
  },
  {
    id: 2,
    name: '阵营 2',
    tileColor: '#3b82f6', // 普通地块染色（蓝色）
    unitColor: '#1d4ed8'
  }
]);

const teams = ref([]);
const currentPlayerId = ref(1);
const selectedTeamId = ref(null);
const actionMode = ref('march'); // 'march' 或 'paint'

// ======= 行军动画和染色倒计时 =======
const teamAnimations = ref({}); // 存储每个队伍的行军路径和当前动画状态
const paintingTasks = ref({}); // 存储染色任务 { hexIndex: { startTime, duration, playerId, teamId } }
const PAINT_DURATION = 60000; // 染色时间：1分钟（60000毫秒）

// ======= 城池数据 =======
const cities = ref([]);
const CITY_NEUTRAL_COLOR = '#facc15'; // 黄色
const CITY_COLORS_BY_PLAYER = {
  1: '#ef4444', // 阵营 1 占领 -> 红色
  2: '#ffffff' // 阵营 2 占领 -> 白色
};

// ======= 阵营初始区域 =======
const spawnAreas = ref([]);

// 每个用户进入场景地图的时候都需要生成一个唯一标识 包含scene_id, 日期, 阵营（每次进入随机）
// 这一部分可以在后续版本中按你的需求增加，这里先专注于基础玩法。

// 计算尖顶六边形坐标
const calculateHexPoints = (x, y) => {
  // 尖顶六边形顶点计算（顶部顶点朝上）
  const points = [];
  for (let i = 0; i < 6; i++) {
    const angle = (Math.PI / 3) * i - Math.PI / 6;
    points.push(
      `${x + hexagonSize.value * Math.cos(angle)},` +
        `${y + hexagonSize.value * Math.sin(angle)}`
    );
  }
  return points.join(' ');
};

// 生成网格
const generateGrid = () => {
  // 六边形间距计算（去除缝隙）
  // 对于尖顶六边形（pointy-top hex）：
  // 宽度 = sqrt(3) * size
  // 高度 = 2 * size
  const hexWidth = hexagonSize.value * Math.sqrt(3); // 六边形的宽度
  const hexSpacingX = hexWidth; // 列间距（等于宽度）
  const hexSpacingY = hexagonSize.value * 1.5; // 行间距（高度 * 0.75）
  const oddRowOffset = hexWidth * 0.5; // 奇数行偏移（宽度的一半）

  const rows = Math.floor(gridSize / hexSpacingY);
  const cols = Math.floor(gridSize / hexSpacingX) + 1;

  rowsCount.value = rows;
  colsCount.value = cols;

  hexagons.value = [];
  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      // 交错排列处理（odd-r 布局）
      const x = col * hexSpacingX + (row % 2) * oddRowOffset;
      const y = row * hexSpacingY;

      if (x < gridSize && y < gridSize) {
        const index = hexagons.value.length;
        hexagons.value.push({
          points: calculateHexPoints(x, y),
          color: '#f0f0f0',
          row,
          col,
          index,
          centerX: x,
          centerY: y,
          type: 'normal', // normal 或 city
          cityId: null,
          isCityCenter: false,
          isPainted: false // 标记是否被染色
        });
      }
    }
  }
};

// 通过 row / col 找到对应 hex
const hexByRowCol = (row, col) => {
  return hexagons.value.find((h) => h.row === row && h.col === col) || null;
};

// odd-r 尖顶六边形网格的邻居偏移
const neighborOffsets = (row) => {
  const isOdd = row % 2 === 1;
  if (isOdd) {
    // odd row
    return [
      [-1, 0],
      [-1, 1],
      [0, -1],
      [0, 1],
      [1, 0],
      [1, 1]
    ];
  } else {
    // even row
    return [
      [-1, 0],
      [-1, -1],
      [0, -1],
      [0, 1],
      [1, 0],
      [1, -1]
    ];
  }
};

// 获取某个 hex 周围一圈六个邻居
const getRingNeighbors = (centerHex) => {
  const result = [];
  const offsets = neighborOffsets(centerHex.row);
  for (const [dr, dc] of offsets) {
    const nr = centerHex.row + dr;
    const nc = centerHex.col + dc;
    if (nr < 0 || nc < 0 || nr >= rowsCount.value || nc >= colsCount.value) {
      continue;
    }
    const neighbor = hexByRowCol(nr, nc);
    if (neighbor) {
      result.push(neighbor);
    }
  }
  return result;
};

// 初始化城池：1 个城池，占 7 个格子（中心 + 外围一圈 6 个）
const initCities = () => {
  if (hexagons.value.length === 0) return;

  // 选取地图中间位置作为首个城池中心
  const centerRow = Math.floor(rowsCount.value / 2);
  const centerCol = Math.floor(colsCount.value / 2);
  const centerHex = hexByRowCol(centerRow, centerCol);
  if (!centerHex) return;

  const neighbors = getRingNeighbors(centerHex);
  const cityHexes = [centerHex, ...neighbors];

  const cityId = 1;
  cities.value.push({
    id: cityId,
    centerIndex: centerHex.index,
    hexIndices: cityHexes.map((h) => h.index),
    owner: 0 // 0 表示未被占领
  });

  // 标记并上色城池格子为黄色
  for (const h of cityHexes) {
    const target = hexagons.value[h.index];
    target.type = 'city';
    target.cityId = cityId;
    target.isCityCenter = h.index === centerHex.index;
    target.color = CITY_NEUTRAL_COLOR;
  }
};

// 根据当前城池状态刷新城池格子的颜色
const refreshCityColors = () => {
  for (const city of cities.value) {
    const color =
      city.owner === 0
        ? CITY_NEUTRAL_COLOR
        : CITY_COLORS_BY_PLAYER[city.owner] || CITY_NEUTRAL_COLOR;
    for (const idx of city.hexIndices) {
      if (hexagons.value[idx]) {
        hexagons.value[idx].color = color;
      }
    }
  }
};

// 设置城池占领者
const setCityOwner = (cityId, playerId) => {
  const city = cities.value.find((c) => c.id === cityId);
  if (!city) return;
  city.owner = playerId;
  refreshCityColors();
};

// 初始化队伍（每个阵营最多 5 支，不同兵种各一支）
const initTeams = () => {
  teams.value = [];
  let idCounter = 1;
  for (const player of players.value) {
    for (let i = 0; i < maxTeamsPerPlayer && i < teamTypes.length; i++) {
      teams.value.push({
        id: idCounter++,
        playerId: player.id,
        type: teamTypes[i],
        hexIndex: null // 还未放到地图上
      });
    }
  }
};

// ======= 计算属性与工具函数 =======
const playerById = (id) =>
  players.value.find((p) => p.id === id) || null;

const currentPlayerTeams = computed(() =>
  teams.value.filter((t) => t.playerId === currentPlayerId.value)
);

const selectedTeam = computed(() =>
  teams.value.find((t) => t.id === selectedTeamId.value) || null
);

const teamLabel = (team) => {
  if (team.hexIndex === null) return '当前未部署在地图上';
  const h = hexagons.value[team.hexIndex];
  if (!h) return '位置未知';
  if (h.type === 'city') {
    return h.isCityCenter
      ? '位于城池中心'
      : '位于城池区域';
  }
  return `位于第 ${h.row} 行，第 ${h.col} 列格子`;
};

const teamShortType = (type) => {
  // 简写显示在小圆点中
  if (!type) return '';
  return type[0]; // 取首字作为简写（骑 / 盾 / 枪 / 弓 / 投）
};

// 处理队伍选择
const selectTeam = (team) => {
  selectedTeamId.value = team.id;
};

// 检查格子是否被染色（不是默认颜色）
const isHexPainted = (hex) => {
  return hex.color !== '#f0f0f0' && hex.color !== CITY_NEUTRAL_COLOR;
};

// 检查格子颜色是否与玩家阵营颜色相同
const isHexColorMatchPlayer = (hex, player) => {
  if (!hex || !player) return false;
  return hex.color === player.tileColor;
};

// 检查是否可以从当前位置染色目标格子
const canPaintFromCurrentPosition = (targetHex) => {
  if (!selectedTeam.value || selectedTeam.value.hexIndex === null) {
    return false;
  }
  
  const currentHex = hexagons.value[selectedTeam.value.hexIndex];
  if (!currentHex) return false;
  
  const player = playerById(selectedTeam.value.playerId);
  if (!player) return false;
  
  // 只能从同颜色的格子（自己阵营的格子）染色相邻的格子
  if (isHexColorMatchPlayer(currentHex, player)) {
    // 检查目标格子是否与当前位置相邻
    const neighbors = getRingNeighbors(currentHex);
    return neighbors.some(n => n.index === targetHex.index);
  }
  
  return false;
};

// 使用BFS查找从起点到终点的路径
const findPath = (startHex, targetHex) => {
  if (!startHex || !targetHex || startHex.index === targetHex.index) {
    return [targetHex.index];
  }

  const queue = [[startHex.index]];
  const visited = new Set([startHex.index]);

  while (queue.length > 0) {
    const path = queue.shift();
    const currentIndex = path[path.length - 1];
    const currentHex = hexagons.value[currentIndex];

    if (!currentHex) continue;

    const neighbors = getRingNeighbors(currentHex);
    for (const neighbor of neighbors) {
      if (neighbor.index === targetHex.index) {
        return [...path, neighbor.index];
      }

      if (!visited.has(neighbor.index)) {
        visited.add(neighbor.index);
        queue.push([...path, neighbor.index]);
      }
    }
  }

  return [startHex.index, targetHex.index]; // 如果找不到路径，直接返回起点和终点
};

// 执行队伍行军动画
const animateTeamMarch = async (team, targetHexIndex) => {
  const teamId = team.id;
  const startHexIndex = team.hexIndex;
  
  if (startHexIndex === null) {
    // 如果队伍还没部署，直接设置位置
    team.hexIndex = targetHexIndex;
    return;
  }

  const startHex = hexagons.value[startHexIndex];
  const targetHex = hexagons.value[targetHexIndex];
  
  if (!startHex || !targetHex) return;

  // 查找路径
  const path = findPath(startHex, targetHex);
  
  if (path.length <= 1) {
    team.hexIndex = targetHexIndex;
    return;
  }

  // 设置动画状态
  const animationStartTime = Date.now();
  teamAnimations.value[teamId] = {
    path: path,
    currentStep: 0,
    isAnimating: true,
    startTime: animationStartTime
  };

  // 逐步移动
  for (let i = 1; i < path.length; i++) {
    const stepStartTime = Date.now();
    teamAnimations.value[teamId].currentStep = i - 1;
    teamAnimations.value[teamId].stepStartTime = stepStartTime;
    
    await new Promise(resolve => setTimeout(resolve, 1000)); // 每个格子1秒
    
    team.hexIndex = path[i];
    if (teamAnimations.value[teamId]) {
      teamAnimations.value[teamId].currentStep = i;
    }
  }

  // 动画完成
  delete teamAnimations.value[teamId];
};

// 开始染色任务
const startPaintingTask = (hexIndex, playerId, teamId) => {
  const taskKey = `${hexIndex}_${teamId}`;
  paintingTasks.value[taskKey] = {
    hexIndex,
    playerId,
    teamId,
    startTime: Date.now(),
    duration: PAINT_DURATION
  };

  // 设置定时器，1分钟后完成染色
  setTimeout(() => {
    const task = paintingTasks.value[taskKey];
    if (task) {
      const hex = hexagons.value[hexIndex];
      const player = playerById(playerId);
      
      if (hex && player && hex.type === 'normal') {
        hexagons.value[hexIndex].color = player.tileColor;
        hexagons.value[hexIndex].isPainted = true;
      }

      // 如果点击的是城池中心格子，则占领城池
      if (hex && hex.type === 'city' && hex.isCityCenter && hex.cityId != null) {
        setCityOwner(hex.cityId, playerId);
      }

      delete paintingTasks.value[taskKey];
    }
  }, PAINT_DURATION);
};

// 获取染色倒计时（秒）
const getPaintingCountdown = (taskKey) => {
  const task = paintingTasks.value[taskKey];
  if (!task) return null;
  
  const elapsed = Date.now() - task.startTime;
  const remaining = Math.max(0, task.duration - elapsed);
  return Math.ceil(remaining / 1000);
};

// 获取队伍显示位置（支持动画插值）
const getTeamDisplayX = (team) => {
  const animation = teamAnimations.value[team.id];
  if (!animation || team.hexIndex === null) {
    const hex = hexagons.value[team.hexIndex];
    return hex?.centerX || 0;
  }

  const currentStep = animation.currentStep;
  const path = animation.path;
  
  if (currentStep < path.length - 1) {
    // 在两个格子之间插值
    const currentHex = hexagons.value[path[currentStep]];
    const nextHex = hexagons.value[path[currentStep + 1]];
    
    if (currentHex && nextHex && animation.stepStartTime) {
      // 计算动画进度（0-1），基于步骤开始时间
      const elapsed = Date.now() - animation.stepStartTime;
      const progress = Math.min(1, elapsed / 1000);
      return currentHex.centerX + (nextHex.centerX - currentHex.centerX) * progress;
    }
  }
  
  const hex = hexagons.value[team.hexIndex];
  return hex?.centerX || 0;
};

const getTeamDisplayY = (team) => {
  const animation = teamAnimations.value[team.id];
  if (!animation || team.hexIndex === null) {
    const hex = hexagons.value[team.hexIndex];
    return hex?.centerY || 0;
  }

  const currentStep = animation.currentStep;
  const path = animation.path;
  
  if (currentStep < path.length - 1) {
    // 在两个格子之间插值
    const currentHex = hexagons.value[path[currentStep]];
    const nextHex = hexagons.value[path[currentStep + 1]];
    
    if (currentHex && nextHex && animation.stepStartTime) {
      // 计算动画进度（0-1），基于步骤开始时间
      const elapsed = Date.now() - animation.stepStartTime;
      const progress = Math.min(1, elapsed / 1000);
      return currentHex.centerY + (nextHex.centerY - currentHex.centerY) * progress;
    }
  }
  
  const hex = hexagons.value[team.hexIndex];
  return hex?.centerY || 0;
};

// 倒计时更新定时器
const countdownTimer = ref(null);
const animationTimer = ref(null);
const updateCountdown = () => {
  // 触发响应式更新
  if (Object.keys(paintingTasks.value).length > 0) {
    // 强制更新
    paintingTasks.value = { ...paintingTasks.value };
  }
};

// 更新动画显示
const updateAnimations = () => {
  // 触发响应式更新以刷新队伍位置
  if (Object.keys(teamAnimations.value).length > 0) {
    teamAnimations.value = { ...teamAnimations.value };
  }
};

onMounted(() => {
  // 根据设备类型调整格子大小
  adjustHexagonSizeForDevice();

  generateGrid();
  initCities();
  initTeams();
  initSpawnAreas();
  updateViewBox();

  // 启动倒计时更新定时器（每秒更新一次）
  countdownTimer.value = setInterval(updateCountdown, 1000);

  // 启动动画更新定时器（每100ms更新一次，使动画更流畅）
  animationTimer.value = setInterval(updateAnimations, 100);

  // 连接WebSocket
  // connectWebSocket();
});

// 组件卸载时清理定时器和WebSocket连接
onUnmounted(() => {
  if (countdownTimer.value) {
    clearInterval(countdownTimer.value);
  }
  if (animationTimer.value) {
    clearInterval(animationTimer.value);
  }

  // 断开WebSocket连接
  // disconnectWebSocket();
});

// 处理六边形点击：移动队伍 + 染色
const handleHexClick = async (e, hex) => {
  e.stopPropagation();

  // 如果正在拖动或拖动距离超过阈值，不处理点击
  // 对于触摸事件使用更大的阈值，因为触摸可能有轻微移动
  const isTouchEvent = e.type && e.type.startsWith('touch');
  const threshold = isTouchEvent ? 20 : 5; // 触摸事件使用更大的阈值
  if (isDragging.value || dragDistance.value > threshold) {
    return;
  }

  if (!selectedTeam.value) {
    return;
  }

  const player = playerById(selectedTeam.value.playerId);
  if (!player) return;

  // 如果队伍正在动画中，不允许新的操作
  if (teamAnimations.value[selectedTeam.value.id]?.isAnimating) {
    return;
  }

  const targetHexIndex = hex.index;

  // 如果是染色模式，先检查是否可以从当前位置染色
  let shouldPaint = false;
  if (actionMode.value === 'paint') {
    shouldPaint = canPaintFromCurrentPosition(hex);
  }

  // 执行行军动画
  const fromHexIndex = selectedTeam.value.hexIndex;
  await animateTeamMarch(selectedTeam.value, targetHexIndex);

  // 发送队伍移动消息
  
  if (fromHexIndex !== null) {
    sendTeamMoveMessage(selectedTeam.value.id, fromHexIndex, targetHexIndex);
  }
  

  // 如果满足染色条件，则开始染色任务
  if (shouldPaint) {
    startPaintingTask(targetHexIndex, player.id, selectedTeam.value.id);
    // 发送染色消息
    // sendPaintMessage(targetHexIndex, player.id, selectedTeam.value.id);
  }
};

// 处理空白区域点击
const handleGridClick = () => {
  // 这里暂时不做处理，保留接口方便后续扩展
};

// 计算绘制城池轮廓的多边形点串（包裹整座城池）
const cityOutlinePoints = (city) => {
  if (!city || !city.hexIndices || city.hexIndices.length === 0) return '';
  const pts = city.hexIndices
    .map((idx) => {
      const h = hexagons.value[idx];
      if (!h) return null;
      return [h.centerX, h.centerY];
    })
    .filter(Boolean);

  if (pts.length === 0) return '';

  // 简单按角度排序，形成一个大致包裹城池的多边形
  const cx =
    pts.reduce((sum, p) => sum + p[0], 0) / pts.length;
  const cy =
    pts.reduce((sum, p) => sum + p[1], 0) / pts.length;

  const sorted = pts.sort((a, b) => {
    const angleA = Math.atan2(a[1] - cy, a[0] - cx);
    const angleB = Math.atan2(b[1] - cy, b[0] - cx);
    return angleA - angleB;
  });

  return sorted.map((p) => `${p[0]},${p[1]}`).join(' ');
};

// ======= 地图拖动与缩放处理 =======
const handleMouseDown = (e) => {
  if (e.button === 0) { // 左键
    isDragging.value = true;
    dragStart.value = { x: e.clientX, y: e.clientY };
    dragDistance.value = 0;
    if (mapContainer.value) {
      mapContainer.value.style.cursor = 'grabbing';
    }
  }
};

const handleMouseMove = (e) => {
  if (isDragging.value) {
    const dx = e.clientX - dragStart.value.x;
    const dy = e.clientY - dragStart.value.y;
    dragDistance.value += Math.sqrt(dx * dx + dy * dy);

    viewBoxOffset.value.x -= dx;
    viewBoxOffset.value.y -= dy;

    updateViewBox();

    dragStart.value = { x: e.clientX, y: e.clientY };
  }
};

const handleMouseUp = () => {
  isDragging.value = false;
  // 延迟重置拖动距离，避免立即触发点击
  setTimeout(() => {
    dragDistance.value = 0;
  }, 100);
  if (mapContainer.value) {
    mapContainer.value.style.cursor = 'default';
  }
};

// 触摸事件处理函数（兼容手机）
const handleTouchStart = (e) => {
  // 只处理单指触摸
  if (e.touches.length === 1) {
    const touch = e.touches[0];
    isDragging.value = true;
    dragStart.value = { x: touch.clientX, y: touch.clientY };
    dragDistance.value = 0;
    // 先不阻止默认行为，等待确认是否为拖动
  }
};

const handleTouchMove = (e) => {
  // 只处理单指触摸
  if (isDragging.value && e.touches.length === 1) {
    const touch = e.touches[0];
    const dx = touch.clientX - dragStart.value.x;
    const dy = touch.clientY - dragStart.value.y;
    const currentDistance = Math.sqrt(dx * dx + dy * dy);

    // 如果移动距离超过阈值，确认为拖动操作
    if (currentDistance > 10) {
      dragDistance.value = currentDistance;

      viewBoxOffset.value.x -= dx;
      viewBoxOffset.value.y -= dy;

      updateViewBox();

      dragStart.value = { x: touch.clientX, y: touch.clientY };

      // 阻止默认行为，防止页面滚动（只有在事件可取消时）
      if (e.cancelable) {
        e.preventDefault();
      }
    }
  }
};

const handleTouchEnd = (e) => {
  isDragging.value = false;
  // 延迟重置拖动距离，避免立即触发点击
  setTimeout(() => {
    dragDistance.value = 0;
  }, 100);
  // 移除触摸后的默认行为处理
};

// 六边形触摸事件处理（用于点击）
let touchStartHex = null;
let touchStartTime = 0;

const handleHexTouchStart = (e, hex) => {
  touchStartHex = hex;
  touchStartTime = Date.now();
  // 不阻止默认行为，允许可能的点击事件
};

const handleHexTouchEnd = async (e, hex) => {
  // 如果是同一个六边形，且触摸时间短（认为是点击而非拖动）
  if (touchStartHex === hex && (Date.now() - touchStartTime) < 300 && dragDistance.value < 10) {
    // 模拟点击事件
    await handleHexClick(e, hex);
  }
  touchStartHex = null;
};


const updateViewBox = () => {
  // 允许地图拖动到边缘甚至稍微超出，让边缘的格子完全可见
  const margin = svgWidth.value * 0.2; // 允许超出20%的边距
  const vx = Math.max(-margin, Math.min(gridSize - svgWidth.value + margin, viewBoxOffset.value.x));
  const vy = Math.max(-margin, Math.min(gridSize - svgHeight.value + margin, viewBoxOffset.value.y));
  viewBox.value = `${vx} ${vy} ${svgWidth.value} ${svgHeight.value}`;
};

// ======= 初始化阵营初始区域 =======
const initSpawnAreas = () => {
  if (hexagons.value.length === 0) return;
  
  const spawnSize = 10; // 10x10 区域
  
  // 阵营1：左下角区域（红色虚线），10x10
  const player1StartRow = Math.max(0, rowsCount.value - spawnSize);
  const player1EndRow = rowsCount.value - 1;
  const player1StartCol = 0;
  const player1EndCol = Math.min(colsCount.value - 1, spawnSize - 1);
  
  const player1Hexes = [];
  const player1 = playerById(1);
  
  for (let row = player1StartRow; row <= player1EndRow && row < rowsCount.value; row++) {
    for (let col = player1StartCol; col <= player1EndCol && col < colsCount.value; col++) {
      const hex = hexByRowCol(row, col);
      if (hex && hex.type === 'normal') {
        player1Hexes.push(hex);
        // 自动染色为阵营1的颜色
        hexagons.value[hex.index].color = player1.tileColor;
        hexagons.value[hex.index].isPainted = true;
      }
    }
  }
  
  // 恢复出生点位边框显示
  spawnAreas.value.push({
    playerId: 1,
    hexIndices: player1Hexes.map(h => h.index),
    strokeColor: '#ef4444', // 红色
    fillColor: '#ef4444'
  });
  
  // 阵营2：右上角区域（蓝色虚线），10x10
  const player2StartRow = 0;
  const player2EndRow = Math.min(rowsCount.value - 1, spawnSize - 1);
  const player2StartCol = Math.max(0, colsCount.value - spawnSize);
  const player2EndCol = colsCount.value - 1;
  
  const player2Hexes = [];
  const player2 = playerById(2);
  
  for (let row = player2StartRow; row <= player2EndRow && row < rowsCount.value; row++) {
    for (let col = player2StartCol; col <= player2EndCol && col < colsCount.value; col++) {
      const hex = hexByRowCol(row, col);
      if (hex && hex.type === 'normal') {
        player2Hexes.push(hex);
        // 自动染色为阵营2的颜色
        hexagons.value[hex.index].color = player2.tileColor;
        hexagons.value[hex.index].isPainted = true;
      }
    }
  }
  
  // 恢复出生点位边框显示
  spawnAreas.value.push({
    playerId: 2,
    hexIndices: player2Hexes.map(h => h.index),
    strokeColor: '#3b82f6', // 蓝色
    fillColor: '#3b82f6'
  });
};

// 计算阵营初始区域轮廓

const spawnAreaOutlinePoints = (spawn) => {
  if (!spawn || !spawn.hexIndices || spawn.hexIndices.length === 0) return '';

  // 获取所有六边形的中心点和顶点
  const centers = [];
  const allPoints = [];

  for (const idx of spawn.hexIndices) {
    const hex = hexagons.value[idx];
    if (!hex) continue;

    centers.push([hex.centerX, hex.centerY]);

    // 解析六边形的顶点
    const points = hex.points.split(' ').map(p => {
      const [x, y] = p.split(',').map(Number);
      return [x, y];
    });
    allPoints.push(...points);
  }

  if (allPoints.length === 0) return '';

  // 计算边界框
  const minX = Math.min(...allPoints.map(p => p[0]));
  const maxX = Math.max(...allPoints.map(p => p[0]));
  const minY = Math.min(...allPoints.map(p => p[1]));
  const maxY = Math.max(...allPoints.map(p => p[1]));

  // 计算中心点
  const cx = centers.reduce((sum, p) => sum + p[0], 0) / centers.length;
  const cy = centers.reduce((sum, p) => sum + p[1], 0) / centers.length;

  // 获取所有唯一的边界点（去重）
  const uniquePoints = [];
  const seen = new Set();
  for (const p of allPoints) {
    const key = `${Math.round(p[0])},${Math.round(p[1])}`;
    if (!seen.has(key)) {
      seen.add(key);
      uniquePoints.push(p);
    }
  }

  // 按角度排序，形成凸包轮廓
  const sorted = uniquePoints.sort((a, b) => {
    const angleA = Math.atan2(a[1] - cy, a[0] - cx);
    const angleB = Math.atan2(b[1] - cy, b[0] - cx);
    return angleA - angleB;
  });

  // 如果点数足够，返回排序后的点；否则使用边界框
  if (sorted.length >= 3) {
    return sorted.map((p) => `${p[0]},${p[1]}`).join(' ');
  } else {
    const margin = hexagonSize.value * 0.5;
    return [
      [minX - margin, minY - margin],
      [maxX + margin, minY - margin],
      [maxX + margin, maxY + margin],
      [minX - margin, maxY + margin]
    ].map(p => `${p[0]},${p[1]}`).join(' ');
  }
};


</script>

<style scoped>
.scene-container {
  display: flex;
  width: 100%;
  height: 100vh;
  box-sizing: border-box;
}

.sidebar {
  width: 280px;
  padding: 16px;
  box-sizing: border-box;
  border-right: 1px solid #ddd;
  background: #f9fafb;
  overflow-y: auto;
}

.sidebar h2 {
  margin: 0 0 12px;
  font-size: 18px;
}

.sidebar h3 {
  margin: 16px 0 8px;
  font-size: 16px;
}

.player-switch {
  display: flex;
  gap: 8px;
  margin-bottom: 12px;
}

.player-tab {
  flex: 1;
  padding: 6px 8px;
  text-align: center;
  border-radius: 6px;
  border: 1px solid #d1d5db;
  cursor: pointer;
  font-size: 14px;
  background: #ffffff;
  user-select: none;
}

.player-tab.active {
  border-color: #2563eb;
  background: #dbeafe;
  color: #1d4ed8;
  font-weight: 600;
}

.action-mode {
  margin-top: 16px;
  margin-bottom: 12px;
}

.action-mode h3 {
  margin: 0 0 8px;
  font-size: 14px;
  font-weight: 600;
}

.mode-switch {
  display: flex;
  gap: 8px;
}

.mode-tab {
  flex: 1;
  padding: 6px 8px;
  text-align: center;
  border-radius: 6px;
  border: 1px solid #d1d5db;
  cursor: pointer;
  font-size: 13px;
  background: #ffffff;
  user-select: none;
  transition: all 0.15s;
}

.mode-tab:hover {
  background: #f3f4f6;
}

.mode-tab.active {
  border-color: #22c55e;
  background: #dcfce7;
  color: #16a34a;
  font-weight: 600;
}

.team-list {
  margin-top: 4px;
}

.team-item {
  border-radius: 8px;
  border: 1px solid #e5e7eb;
  padding: 8px 10px;
  margin-bottom: 8px;
  background: #ffffff;
  cursor: pointer;
  transition: background 0.15s, border-color 0.15s, box-shadow 0.15s;
  font-size: 13px;
}

.team-item:hover {
  background: #f3f4f6;
}

.team-item.selected {
  border-color: #22c55e;
  box-shadow: 0 0 0 1px #22c55e33;
}

.team-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 2px;
}

.team-type {
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 6px;
}

.team-avatar {
  font-size: 18px;
}

.team-status {
  font-size: 12px;
  color: #6b7280;
}

.team-sub {
  font-size: 12px;
  color: #4b5563;
}

.tips {
  margin-top: 16px;
  font-size: 12px;
  color: #4b5563;
}

.tips ul {
  padding-left: 18px;
  margin: 4px 0 0;
}

.tips li {
  margin-bottom: 2px;
}

.tips li.indent {
  margin-left: 12px;
}

.hex-grid {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  background: #e5e7eb;
  cursor: grab;
  position: relative;
}

.hex-grid:active {
  cursor: grabbing;
}

.hex-grid svg {
  background: #f9fafb;
  display: block;
}
</style>