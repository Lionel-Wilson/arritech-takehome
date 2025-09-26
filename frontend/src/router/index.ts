import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import UsersPage from "../views/UsersPage.vue"
import CreateUserPage from "../views/CreateUserPage.vue"
import EditUserPage from "../views/EditUserPage.vue"

const routes: RouteRecordRaw[] = [
  { path: '/', redirect: '/users' },
  { path: '/users', component: UsersPage },
  { path: '/users/new', component: CreateUserPage },
  { path: '/users/:id', component: EditUserPage, props: true },
]

export default createRouter({
  history: createWebHistory(), // HTML5 history mode
  routes,
})
