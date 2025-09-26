<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { api } from '../api'
import type { User } from '../types'
import UserForm from '../components/UserForm.vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const user = ref<User | null>(null)

onMounted(async () => {
  const { data } = await api.get(`/users/${route.params.id}`)
  user.value = data.user || data
})
</script>

<template>
  <div class="section">
    <el-card class="panel el-card--always-shadow">
  <el-skeleton v-if="!user" animated />
  <UserForm v-else mode="edit" :user="user!" @submitted="() => router.push('/users')" />
  </el-card>
  </div>
</template>

<style scoped>
.panel :deep(.el-card__body) { padding: 16px; }
</style>
