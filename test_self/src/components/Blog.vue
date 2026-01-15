<script setup lang="ts">
import { activeSection } from '@/main';
import { ref } from 'vue';

const blogPosts = ref([
  {
      id: 1,
      title: 'Vue 3 Composition API 使用指南',
      excerpt: '深入学习Vue 3的Composition API，了解如何组织你的代码逻辑。',
      date: '2023-06-15',
      image: 'https://images.unsplash.com/photo-1581276879432-15e50529f34b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80'
  },
  {
      id: 2,
      title: 'CSS Grid布局实战',
      excerpt: '掌握CSS Grid布局，创建复杂而响应式的网页布局。',
      date: '2023-05-28',
      image: 'https://images.unsplash.com/photo-1581276879432-15e50529f34b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80'
  },
  {
      id: 3,
      title: '前端性能优化技巧',
      excerpt: '提高网站加载速度和运行效率的实用技巧和最佳实践。',
      date: '2023-04-10',
      image: 'https://images.unsplash.com/photo-1533750349088-cd871a92f312?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1000&q=80'
  }
]);
</script>

<template>
  <section id="blog" class="section" :class="{active: activeSection === 'blog'}">
      <h2 class="section-title">技术博客</h2>
      <div class="blog-posts">
          <article class="blog-card" v-for="post in blogPosts" :key="post.id">
              <img :src="post.image" class="blog-img" :alt="post.title">
              <div class="blog-content">
                  <div class="blog-date">{{ post.date }}</div>
                  <h3 class="blog-title">{{ post.title }}</h3>
                  <p class="blog-excerpt">{{ post.excerpt }}</p>
                  <a href="#" class="btn">阅读更多</a>
              </div>
          </article>
      </div>
  </section>
</template>

<style scoped>
/* 与其他页面统一的整体宽度 */
.section {
  width: 100%;
  max-width: 1080px;
  margin: 0 auto;
  padding: 1.5rem 1rem 3rem;
  box-sizing: border-box;
}

.section-title {
  font-size: 1.6rem;
  margin: 0 0 1.25rem;
  padding-bottom: 0.5rem;
  border-bottom: 2px solid var(--accent-color);
  color: var(--primary-color);
}

/* 博客样式 */
.blog-posts {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 1.75rem;
  margin-top: 1.5rem;
}

.blog-card {
  background: #ffffff;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.08);
  transition: var(--transition);
}

.blog-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.blog-img {
  height: 180px;
  width: 100%;
  object-fit: cover;
}

.blog-content {
  padding: 1.5rem;
}

.blog-date {
  color: var(--text-light);
  font-size: 0.85rem;
  margin-bottom: 0.5rem;
}

.blog-title {
  font-size: 1.2rem;
  margin-bottom: 0.75rem;
  color: var(--primary-color);
}

.blog-excerpt {
  color: var(--text-light);
  margin-bottom: 1.25rem;
  font-size: 0.95rem;
  line-height: 1.7;
}

.btn {
  display: inline-block;
  background: var(--accent-color);
  color: #ffffff;
  padding: 0.6rem 1.2rem;
  border-radius: 999px;
  text-decoration: none;
  font-size: 0.9rem;
  font-weight: 500;
  transition: var(--transition);
}

.btn:hover {
  background: #ff5252;
}

/* 手机端优化 */
@media (max-width: 768px) {
  .section {
    padding: 1rem 0.75rem 2rem;
  }

  .blog-posts {
    grid-template-columns: 1fr;
    gap: 1.25rem;
  }

  .blog-content {
    padding: 1.25rem 1rem;
  }
}
</style>
