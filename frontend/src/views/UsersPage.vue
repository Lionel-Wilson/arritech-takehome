<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { api } from '@/api'          // adjust path if you don't use '@'
import type { User } from '@/types'
import { ElMessageBox, ElMessage } from 'element-plus'

const users = ref<User[]>([])          // never null
const loading = ref(false)
const search = ref('')                  // safe default
const page = ref(1)                     // safe default
const pageSize = ref(10)                // safe default
const total = ref(0)                    // safe default

async function fetchUsers() {
  loading.value = true
  try {
    const { data } = await api.get('/users/', {
      params: {
        q: search.value || '',
        page: page.value || 1,
        page_size: pageSize.value || 10,
      }
    })
    // make sure these are always defined as arrays/numbers
    users.value = Array.isArray(data?.users) ? data.users : []
    total.value = Number.isFinite(data?.total) ? data.total : users.value.length
  } catch (e) {
    ElMessage.error('Failed to load users')
    users.value = []    // <- never null
    total.value = 0
  } finally {
    loading.value = false
  }
}

// explicit actions (avoid re-render storms while typing)
function onSearchClick() {
  page.value = 1  // reset to 1 when searching
  fetchUsers()
}

// keep these watchers very small/safe
watch(page, fetchUsers)
watch(pageSize, () => {
  page.value = 1
  fetchUsers()
})

async function confirmDelete(row: User) {
  try {
    await ElMessageBox.confirm(`Delete ${row.firstname} ${row.lastname}?`, 'Confirm', { type: 'warning' })
    await api.delete(`/users/${row.id}`)
    ElMessage.success('User deleted')
    fetchUsers()
  } catch (_) {}
}

onMounted(fetchUsers)
</script>

<template>
  <div class="section">
    <el-card class="panel el-card--always-shadow">
      <div class="toolbar">
        <el-input
          v-model="search"
          placeholder="Search name or email"
          clearable
          @keyup.enter="onSearchClick"
        />
        <el-button type="primary" @click="onSearchClick">Search</el-button>
        <router-link to="/users/new">
          <el-button>New User</el-button>
        </router-link>
      </div>

      <el-table :data="users" v-loading="loading" stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="firstname" label="First name" />
        <el-table-column prop="lastname" label="Last name" />
        <el-table-column prop="email" label="Email" />
        <el-table-column prop="age" label="Age" width="80" />
        <el-table-column label="Actions" width="180">
          <template #default="{ row }">
            <router-link :to="`/users/${row.id}`">
              <el-button size="small">Edit</el-button>
            </router-link>
            <el-button size="small" type="danger" @click="confirmDelete(row)">Delete</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pager">
        <el-pagination
          v-model:current-page="page"
          v-model:page-size="pageSize"
          :total="total"
          layout="prev, pager, next, sizes, total"
          :page-sizes="[5, 10, 20, 50]"
        />
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.section { display: grid; place-items: start center; }
.panel   { width: 100%; background: #fff; }
.panel :deep(.el-card__body) { padding: 16px; }
.toolbar { display: flex; gap: 12px; align-items: center; margin-bottom: 12px; }
.toolbar :deep(.el-input) { max-width: 320px; }
.pager { margin-top: 12px; display: flex; justify-content: flex-end; }
</style>
