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

    <main class="section">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.site-header {
  display: flex;
  flex-direction: column;        /* stack title + nav on small screens */
  align-items: center;
  gap: 6px;
  margin-bottom: 12px;
}

@media (min-width: 640px) {
  .site-header {
    flex-direction: row;         /* side by side on wider screens */
    justify-content: center;     /* centered as a group */
    gap: 16px;
  }
}

.title { margin: 0; font-weight: 600; }
.nav { display: flex; gap: 12px; }
.nav a { color: var(--link); text-decoration: none; font-weight: 500; }
.nav a:hover { text-decoration: underline; }
</style>
