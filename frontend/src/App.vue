<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="container">
    <header class="site-header">
      <h2 class="title">Arritech – Users</h2>
      <nav class="nav">
        <RouterLink to="/users">List</RouterLink>
        <RouterLink to="/users/new">Create</RouterLink>
      </nav>
    </header>

    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.site-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 16px;
}

.title {
  margin: 0;
  font-weight: 600;
}

.nav a {
  color: var(--link);
  text-decoration: none;
  font-weight: 500;
  margin-left: 8px;
}
.nav a:hover { text-decoration: underline; }
</style>
