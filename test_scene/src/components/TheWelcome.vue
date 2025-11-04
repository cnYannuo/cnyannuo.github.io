<template>
  <div class="hex-grid">
    <svg 
      width="4000" 
      height="3500" 
      viewBox="0 0 4000 3500"
      @click="handleGridClick"
    >
      <!-- 生成尖顶六边形网格 -->
      <polygon 
        v-for="(hex, index) in hexagons" 
        :key="index"
        :points="hex.points"
        :fill="hex.color"
        stroke="#333"
        stroke-width="0.5"
        @click="(e) => handleHexClick(e, hex)"
        :data-index="index"
      />
    </svg>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
const gridSize = 4000;
const hexagons = ref([]);
const hexagonSize = 24; // 单个六边形边长


// 每个用户进入场景地图的时候都需要生成一个唯一标识 包含scene_id, 日期, 阵营（每次进入随机）



// 计算尖顶六边形坐标
const calculateHexPoints = (x, y) => {
  // 尖顶六边形顶点计算（顶部顶点朝上）
  const points = [];
  for (let i = 0; i < 6; i++) {
    const angle = (Math.PI / 3) * i - (Math.PI / 6);
    points.push(
      `${x + hexagonSize * Math.cos(angle)},` +
      `${y + hexagonSize * Math.sin(angle)}`
    );
  }
  return points.join(' ');
};

// 生成网格
const generateGrid = () => {
  const rows = Math.floor(gridSize / (hexagonSize * Math.sqrt(3)));
  const cols = Math.floor(gridSize / (hexagonSize * 1.5)) + 1;
  
  hexagons.value = [];
  console.log("rows: "  + rows + ",cols:" + cols);
  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      // 交错排列处理
      const x = col * (hexagonSize * 1.75) + 
                 (row % 2) * (hexagonSize * 0.9);
      const y = row * (hexagonSize * Math.sqrt(3) * 0.89);
      
      if (x < gridSize && y < gridSize) {
        hexagons.value.push({
          points: calculateHexPoints(x, y),
          color: '#f0f0f0',
          row,
          col,
          index: hexagons.value.length
        });
      }
    }
  }
};

// 处理六边形点击
const handleHexClick = (e, hex) => {
  e.stopPropagation();
  const newColor = hex.color === '#f0f0f0' ? '#4ade80' : '#f0f0f0';
  hexagons.value[hex.index].color = newColor;
  // console.log(`点击了第${hex.row}行第${hex.col}列的六边形`);
};

// 处理空白区域点击
const handleGridClick = (e) => {
  console.log('点击了网格空白区域');
};

onMounted(() => {
  generateGrid();
});
</script>

<style scoped>
.hex-grid {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 4000px;
  height: 3500px;
  border: 1px solid #ddd;
}
</style>