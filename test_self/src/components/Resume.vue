<script setup lang="ts">
import { activeSection } from '@/main';
</script>

<template>
  <section id="resume" class="section" :class="{active: activeSection === 'resume'}">
      <h2 class="section-title">我的简历</h2>
      <div class="resume-content">
          <div class="sidebar">
              <h3>个人信息</h3>
              <p><i class="fas fa-user"></i> 张三</p>
              <p><i class="fas fa-envelope"></i> zhangsan@example.com</p>
              <p><i class="fas fa-phone"></i> 138-8888-8888</p>
              <p><i class="fas fa-map-marker-alt"></i> 北京市海淀区</p>
              
              <h3>技能专长</h3>
              <p>Vue.js ⭐⭐⭐⭐⭐</p>
              <p>JavaScript ⭐⭐⭐⭐⭐</p>
              <p>HTML/CSS ⭐⭐⭐⭐⭐</p>
              <p>UI/UX设计 ⭐⭐⭐⭐</p>
          </div>
          
          <div class="main-content">
              <h3>经历工作</h3>
              <div class="experience">
                  <h4>高级前端工程师 - ABC科技有限公司</h4>
                  <p class="date">2020年至今</p>
                  <p>负责公司主要产品的前端架构设计与实现，使用Vue.js开发复杂单页面应用。</p>
                  
                  <h4>前端开发工程师 - XYZ互联网公司</h4>
                  <p class="date">2018-2020年</p>
                  <p>参与多个企业级Web应用开发，负责前端页面实现与性能优化。</p>
              </div>
              
              <h3>教育背景</h3>
              <div class="education">
                  <h4>计算机科学与技术硕士 - 北京大学</h4>
                  <p class="date">2015-2018年</p>
                  
                  <h4>软件工程学士 - 南京大学</h4>
                  <p class="date">2011-2015年</p>
              </div>
          </div>
      </div>
      
      <h3 class="section-title">我的技能</h3>
      <div class="skills-list">
          <div class="skill-item">
              <h3><i class="fas fa-code"></i> 前端开发</h3>
              <p>熟练掌握HTML5, CSS3, JavaScript, Vue.js, React等前端技术</p>
          </div>
          <div class="skill-item">
              <h3><i class="fas fa-paint-brush"></i> UI/UX设计</h3>
              <p>能够设计美观的用户界面并提供良好的用户体验</p>
          </div>
          <div class="skill-item">
              <h3><i class="fas fa-mobile-alt"></i> 响应式设计</h3>
              <p>能够创建在各种设备上都能完美显示的响应式网站</p>
          </div>
          <div class="skill-item">
              <h3><i class="fas fa-server"></i> 后端知识</h3>
              <p>了解Node.js, Express, MongoDB等后端技术</p>
          </div>
      </div>
  </section>
</template>

<style scoped>
.resume-content {
            display: grid;
            grid-template-columns: 1fr 2fr;
            gap: 2rem;
            background: white;
            border-radius: 10px;
            padding: 2rem;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
        }
        
        .sidebar {
            background: var(--primary-color);
            color: black;
            padding: 2rem;
            border-radius: 10px;
        }
        
        .main-content {
            padding: 0 1rem;
        }
        
        .section-title {
            font-size: 1.8rem;
            margin-bottom: 1.5rem;
            padding-bottom: 0.5rem;
            border-bottom: 2px solid var(--accent-color);
            color: var(--primary-color);
        }
        
        .skills-list {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
            gap: 1.5rem;
            margin-top: 2rem;
        }
        
        .skill-item {
            background: white;
            padding: 1.5rem;
            border-radius: 10px;
            box-shadow: 0 3px 10px rgba(0, 0, 0, 0.08);
            transition: var(--transition);
        }
        
        .skill-item:hover {
            transform: translateY(-5px);
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
        }
        
        .skill-item h3 {
            color: var(--primary-color);
            margin-bottom: 1rem;
            display: flex;
            align-items: center;
        }
        
        .skill-item h3 i {
            margin-right: 10px;
            color: var(--accent-color);
        }
</style>
