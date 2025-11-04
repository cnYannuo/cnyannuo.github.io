<script setup lang="ts">
import { setActiveSection,activeSection } from '@/main';

</script>

<template>
  <section id="contact" class="section" :class="{active: activeSection === 'contact'}">
      <h2 class="section-title">联系我</h2>
      <div style="background: white; padding: 2rem; border-radius: 10px; box-shadow: 0 5px 15px rgba(0,0,0,0.05); text-align: center;">
          <h3>感谢您访问我的个人网站</h3>
          <p>如果您有任何问题或合作意向，请通过以下方式联系我：</p>
          <p>邮箱：zhangsan@example.com</p>
          <p>电话：138-8888-8888</p>
          <div style="margin-top: 2rem;">
              <a href="#" class="btn">发送邮件</a>
          </div>
      </div>
  </section>
</template>

<style scoped>

</style>
