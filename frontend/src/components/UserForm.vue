<script setup lang="ts">
import { Form as VForm, Field, ErrorMessage } from 'vee-validate'
import * as yup from 'yup'
import { api } from '../api'
import type { User } from '../types'
import { ElMessage } from 'element-plus'

const props = defineProps<{ mode: 'create' | 'edit'; user?: User }>()
const emit = defineEmits<{ (e: 'submitted'): void }>()

const schema = yup.object({
  firstname: yup.string().required('First name is required'),
  lastname:  yup.string().required('Last name is required'),
  email:     yup.string().email('Invalid email').required('Email is required'),
  age:       yup.number().typeError('Age must be a number').integer().min(18, 'Must be 18+').required(),
  phonenumber:  yup.string().required('Phonenumber is required')
})

const initial = {
  firstname: props.user?.firstname ?? '',
  lastname:  props.user?.lastname ?? '',
  email:     props.user?.email ?? '',
  age:       props.user?.age ?? 18,
  phonenumber: props.user?.phonenumber??''
}

async function onSubmit(values: any) {
  try {
    if (props.mode === 'create') await api.post('/users/', values)
    else await api.patch(`/users/${props.user!.id}`, values)
    ElMessage.success(props.mode === 'create' ? 'User created' : 'User updated')
    emit('submitted')
  } catch (err: any) {
    ElMessage.error(err?.response?.data?.error || 'Request failed')
  }
}
</script>

<template>
  <el-card class="card">
    <!-- key forces the form to re-mount when editing a different user or switching to 'create' -->
    <VForm
      :key="props.user?.id ?? 'create'"
      :initial-values="initial"
      :validation-schema="schema"
      @submit="onSubmit"
    >
      <div class="row">
        <label>First name</label>
        <Field name="firstname" v-slot="{ field }">
          <el-input
            :model-value="field.value"
            @update:modelValue="field.onChange"
            @blur="field.onBlur"
          />
        </Field>
        <ErrorMessage name="firstname" class="err" />
      </div>

      <div class="row">
        <label>Last name</label>
        <Field name="lastname" v-slot="{ field }">
          <el-input
            :model-value="field.value"
            @update:modelValue="field.onChange"
            @blur="field.onBlur"
          />
        </Field>
        <ErrorMessage name="lastname" class="err" />
      </div>

      <div class="row">
        <label>Email</label>
        <Field name="email" v-slot="{ field }">
          <el-input
            :model-value="field.value"
            @update:modelValue="field.onChange"
            @blur="field.onBlur"
          />
        </Field>
        <ErrorMessage name="email" class="err" />
      </div>

      <div class="row">
        <label>Phone Number</label>
        <Field name="phonenumber" v-slot="{ field }">
          <el-input
            :model-value="field.value"
            @update:modelValue="field.onChange"
            @blur="field.onBlur"
          />
        </Field>
        <ErrorMessage name="phonenumber" class="err" />
      </div>

      <div class="row">
        <label>Age</label>
        <Field name="age" v-slot="{ field }">
          <el-input-number
            :model-value="field.value"
            @update:modelValue="field.onChange"
            @blur="field.onBlur"
            :min="0"
          />
        </Field>
        <ErrorMessage name="age" class="err" />
      </div>

      <el-button type="primary" native-type="submit">
        {{ props.mode === 'create' ? 'Create' : 'Save' }}
      </el-button>
    </VForm>
  </el-card>
</template>

<style scoped>
.card { max-width: 620px; margin: 24px auto; }
.row { margin-bottom:12px; display:grid; grid-template-columns:130px 1fr; align-items:center; gap:12px; }
.err { color:#f56c6c; font-size:12px; grid-column:2; }
</style>
