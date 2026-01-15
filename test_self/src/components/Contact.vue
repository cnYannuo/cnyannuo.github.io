<script setup lang="ts">
import { setActiveSection,activeSection } from '@/main';

</script>

<template>
  <section id="contact" class="section" :class="{active: activeSection === 'contact'}">
    <h2 class="section-title">联系我</h2>
    <div class="contact-card">
      <h3>感谢您访问我的个人网站</h3>
      <p class="contact-text">
        如果您有任何问题或合作意向，请通过以下方式联系我：
      </p>
      <p class="contact-info">邮箱：zhangsan@example.com</p>
      <p class="contact-info">电话：138-8888-8888</p>
      <div class="contact-actions">
        <a href="#" class="btn">发送邮件</a>
      </div>
    </div>
  </section>
</template>

<style scoped>
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

.contact-card {
  background: #ffffff;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
  text-align: center;
}

.contact-card h3 {
  margin-bottom: 0.75rem;
  font-size: 1.25rem;
}

.contact-text {
  margin-bottom: 1rem;
  color: var(--text-light);
  font-size: 0.95rem;
}

.contact-info {
  margin: 0.25rem 0;
  font-size: 0.95rem;
}

.contact-actions {
  margin-top: 2rem;
}

.btn {
  display: inline-block;
  background: var(--accent-color);
  color: #ffffff;
  padding: 0.8rem 1.6rem;
  border-radius: 999px;
  text-decoration: none;
  font-weight: 500;
  transition: var(--transition);
}

.btn:hover {
  background: #ff5252;
}

@media (max-width: 768px) {
  .section {
    padding: 1rem 0.75rem 2rem;
  }

  .contact-card {
    padding: 1.5rem 1.25rem;
  }
}
</style>
