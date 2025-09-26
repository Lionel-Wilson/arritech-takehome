<script setup lang="ts">
import UserForm from '../components/UserForm.vue'
import { useRouter } from 'vue-router'
const router = useRouter()
</script>

<template>
  <el-card class="card">
  <UserForm mode="create" @submitted="() => router.push('/users')" />
  </el-card>
</template>

<style scoped>
.card {
  background: #fff;
  max-width: 640px;
  margin: 0 auto;
}
</style>
