<script setup lang="ts">
import { RouterLink, RouterView, useRoute } from 'vue-router'
const route = useRoute()
</script>

<template>
  <div class="container">
    <div class="nav-links">
      <RouterLink
        to="/"
        :class="{ active: route.path === '/' }"
        class="nav-link"
      >首页</RouterLink>
      <RouterLink
        to="/resume"
        :class="{ active: route.path === '/resume' }"
        class="nav-link"
      >简历</RouterLink>
      <RouterLink
        to="/blog"
        :class="{ active: route.path === '/blog' }"
        class="nav-link"
      >博客</RouterLink>
      <RouterLink
        to="/contact"
        :class="{ active: route.path === '/contact' }"
        class="nav-link"
      >联系我</RouterLink>
      <!-- <RouterLink to="/">首页</RouterLink>
      <RouterLink to="/resume">简历</RouterLink> -->
      <!-- <RouterLink to="/blog">博客</RouterLink> -->
      <!-- <RouterLink to="/contact">联系我</RouterLink> -->
      </div>
    <RouterView />
  </div>
</template>

<style scoped>
.container {
            /* max-width: 1200px; */
            margin: 0 auto;
            padding: 0 20px;
        }
.nav-links {
  display: flex;
  list-style: none;
}
.nav-links a {
  color: #2c3e50;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s ease;
  padding: 0.5rem 0.5rem;
  position: relative;
}
.nav-links a:after {
  content: '';
  position: absolute;
  width: 0;
  height: 2px;
  bottom: 0;
  left: 0;
  background-color: var(--accent-color);
  transition: all 0.3s ease;
}
.nav-links a:hover:after {
  width: 100%;
}
.nav-links a.active {
  color: var(--accent-color);
}
.nav-links a.active:after {
  width: 100%;
}

.section-title {
            font-size: 1.8rem;
            margin-bottom: 1.5rem;
            padding-bottom: 0.5rem;
            border-bottom: 2px solid var(--accent-color);
            color: var(--primary-color);
        }
.section {
            display: none;
            animation: fadeIn 0.5s ease;
        }

        .section.active {
            display: block;
        }

/* @media (max-width: 768px) {
  .nav-links {
    display: none;
  }
} */
</style>
