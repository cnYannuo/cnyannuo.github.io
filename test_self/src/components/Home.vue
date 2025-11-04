<script setup lang="ts">
import { setActiveSection,activeSection } from '@/main';

</script>

<template>
  <section id="home" class="section" :class="{active: activeSection === 'home'}">
      <div class="hero">
          <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80" class="hero-img" alt="个人头像">
          <h1>你好，我是张三</h1>
          <p>热爱创造美观且功能强大的Web应用</p>
          <!-- 一名前端开发工程师 & UI设计师， -->
          <a href="#" class="btn" @click="setActiveSection('resume')">查看我的简历</a>
      </div>
  </section>
</template>

<style scoped>
.hero {
            text-align: center;
            padding: 4rem 0;
            background: white;
            border-radius: 10px;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
            margin-bottom: 2rem;
        }
        
        .hero h1 {
            font-size: 3rem;
            margin-bottom: 1rem;
            color: var(--primary-color);
        }
        
        .hero p {
            font-size: 1.2rem;
            color: var(--text-light);
            max-width: 700px;
            margin: 0 auto 2rem;
        }
        
        .hero-img {
            width: 180px;
            height: 180px;
            border-radius: 50%;
            object-fit: cover;
            border: 5px solid var(--secondary-color);
            margin-bottom: 1.5rem;
        }
        
        .btn {
            display: inline-block;
            background: var(--accent-color);
            color: white;
            padding: 0.8rem 1.5rem;
            border-radius: 50px;
            text-decoration: none;
            font-weight: 500;
            transition: var(--transition);
            border: none;
            cursor: pointer;
        }
        
        .btn:hover {
            background: #ff5252;
            transform: translateY(-3px);
            box-shadow: 0 5px 15px rgba(255, 107, 107, 0.3);
        }
</style>
