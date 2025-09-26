<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { api } from '../api'
import type { User } from '../types'
import { useRouter } from 'vue-router'
import {ElMessage, ElMessageBox} from "element-plus";

const router = useRouter()
const users = ref<User[]>([])
const loading = ref(false)
const search = ref('')
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

async function fetchUsers() {
  loading.value = true
  try {
    const { data } = await api.get('/users/', {
      params: { q: search.value, page: page.value, page_size: pageSize.value }
    })
    users.value = data.users || data || []
    total.value = data.total || users.value.length
  } finally { loading.value = false }
}

async function confirmDelete(row: User) {
  try {
    await ElMessageBox.confirm(`Delete user ${row.firstname} ${row.lastname}?`, 'Confirm', { type: 'warning' })
    await api.delete(`/users/${row.id}`)
    ElMessage.success('User deleted')
    fetchUsers()
  } catch (_) {}
}

onMounted(fetchUsers)
watch([search, page, pageSize], fetchUsers)
</script>

<template>
  <el-card class="card">
    <div style="display:flex; gap:12px; align-items:center; margin-bottom:12px">
      <el-input v-model="search" placeholder="Search name or email" clearable style="max-width:320px"/>
      <el-button type="primary" @click="fetchUsers">Search</el-button>
      <el-button @click="router.push('/users/new')">New User</el-button>
    </div>

    <el-table :data="users" v-loading="loading" stripe>
      <el-table-column prop="id" label="ID" width="80"/>
      <el-table-column prop="firstname" label="First name"/>
      <el-table-column prop="lastname" label="Last name"/>
      <el-table-column prop="email" label="Email"/>
      <el-table-column prop="age" label="Age" width="80"/>
      <el-table-column label="Actions" width="180">
        <template #default="{ row }">
          <el-button size="small" @click="router.push(`/users/${row.id}`)">Edit</el-button>
          <el-button size="small" type="danger" @click="confirmDelete(row)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div style="margin-top:12px; display:flex; justify-content:flex-end">
      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :total="total"
        layout="prev, pager, next, sizes, total"
        :page-sizes="[5,10,20,50]"
      />
    </div>
  </el-card>
</template>


<style scoped>
.card {
  background: #fff;
  max-width: 900px;        /* slightly narrower than the container */
  margin: 0 auto;          /* centered inside the container */
}
</style>
