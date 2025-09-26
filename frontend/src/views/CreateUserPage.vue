<script setup lang="ts">
import UserForm from '../components/UserForm.vue'
import { useRouter } from 'vue-router'
const router = useRouter()
</script>

<template>
  <div class="section">
    <el-card class="panel el-card--always-shadow">
  <UserForm mode="create" @submitted="() => router.push('/users')" />
  </el-card>
  </div>
</template>

<style scoped>
.panel :deep(.el-card__body) { padding: 16px; }
</style>
