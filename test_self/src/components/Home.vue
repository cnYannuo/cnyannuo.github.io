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
  /* 保持和简历页一致的整体宽度 */
  .section {
    width: 100%;
    max-width: 1080px;
    margin: 0 auto;
    padding: 1.5rem 1rem 3rem;
    box-sizing: border-box;
  }

  .hero {
    text-align: center;
    padding: 3rem 1.5rem;
    background: #ffffff;
    border-radius: 12px;
    /* box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05); */
    margin-bottom: 2rem;
  }

  .hero h1 {
    font-size: 2.4rem;
    margin-bottom: 1rem;
    color: var(--primary-color);
  }

  .hero p {
    font-size: 1.05rem;
    color: var(--text-light);
    max-width: 640px;
    margin: 0 auto 2rem;
  }

  .hero-img {
    width: 160px;
    height: 160px;
    border-radius: 50%;
    object-fit: cover;
    border: 5px solid var(--secondary-color);
    margin-bottom: 1.5rem;
  }

  .btn {
    display: inline-block;
    background: var(--accent-color);
    color: #ffffff;
    padding: 0.8rem 1.5rem;
    border-radius: 999px;
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

  /* 手机端优化 */
  @media (max-width: 768px) {
    .section {
      padding: 1rem 0.75rem 2rem;
    }

    .hero {
      padding: 2.25rem 1.25rem;
    }

    .hero h1 {
      font-size: 2rem;
    }

    .hero p {
      font-size: 0.95rem;
    }

    .hero-img {
      width: 136px;
      height: 136px;
    }
  }
  </style>
