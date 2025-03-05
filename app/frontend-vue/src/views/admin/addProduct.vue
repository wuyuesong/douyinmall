<template>
  <div class="admin-container">
    <Header></Header>
    <div class="form-wrapper">
      <el-form 
        :model="form" 
        label-width="auto" 
        class="form-container"
        :rules="rules"
        ref="formRef"
      >
        <el-form-item label="商品名称" prop="name">
          <el-input v-model="form.name" />
        </el-form-item>

        <el-form-item label="单价" prop="price">
          <el-input v-model.number="form.price" type="number" />
        </el-form-item>

        <el-form-item label="库存" prop="stock">
          <el-input v-model.number="form.stock" type="number" />
        </el-form-item>

        <el-form-item label="商品描述" prop="desc">
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
        <el-form-item label="商品图片">
          <el-upload
            class="upload-demo"
            :auto-upload="false"
            :on-change="handleFileChange"
            :file-list="fileList"
            list-type="picture"
            accept="image/*"
          >
            <el-button type="primary">点击上传</el-button>
            <template #tip>
              <div class="el-upload__tip">
                请上传JPEG/PNG格式的图片, 大小不超过2MB
              </div>
            </template>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">添加商品</el-button>
          <el-button @click="handleCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Header from '../../components/AdminHeader.vue';
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import type { UploadFile, UploadFiles } from 'element-plus'
import { useRouter } from 'vue-router'
import type { FormInstance } from 'element-plus'

const router = useRouter()

// 表单数据结构
const form = reactive({
  name: '',
  price: '',
  stock: '',
  desc: '',
  images: [] as File[] // 存储图片文件
})

// 文件列表用于显示预览
const fileList = ref<UploadFile[]>([])


const formRef = ref<FormInstance>()

const rules = reactive({
  name: [
    { required: true, message: '请输入商品名称', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在2到50个字符', trigger: 'blur' }
  ],
  price: [
    { required: true, message: '请输入单价', trigger: 'blur' },
    { type: 'number', message: '必须为数字值' },
    { 
      validator: (_: any, value: number, callback: any) => {
        if (value < 0) callback(new Error('单价不能为负数'))
        else callback()
      }, 
      trigger: 'blur' 
    }
  ],
  stock: [
    { required: true, message: '请输入库存', trigger: 'blur' },
    { 
      validator: (_: any, value: number, callback: any) => {
        if (!Number.isInteger(value)) {
          callback(new Error('必须为整数'))
        } else if (value < 0) {
          callback(new Error('库存不能为负数'))
        } else {
          callback()
        }
      }, 
      trigger: 'blur' 
    }
  ],
  desc: [
    { required: true, message: '请输入商品描述', trigger: 'blur' },
    { min: 10, max: 500, message: '长度在10到500个字符', trigger: 'blur' }
  ]
})

// 处理文件选择
const handleFileChange = (file: UploadFile, files: UploadFiles) => {
  const isImage = file.raw?.type.startsWith('image/')
  const isLt2M = file.raw?.size && file.raw.size / 1024 / 1024 < 2


  if (!isImage) {
    ElMessage.error('只能上传图片文件！')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过2MB!')
    return false
  }

  // 将文件存入表单数据
  if (file.raw) {
    form.images.push(file.raw)
  }
  return true
}

const onSubmit = async () => {
  // 构造表单数据
  const formData = new FormData()
  formData.append('name', form.name)
  formData.append('price', form.price)
  formData.append('stock', form.stock)
  formData.append('desc', form.desc)
  
  // 添加图片文件
  form.images.forEach((file, index) => {
    formData.append(`images[${index}]`, file)
  })

  try {
    // 替换为你的API地址
    const response = await fetch(import.meta.env.VITE_API_BASE + '/products', {
      method: 'POST',
      headers: {
        Authorization: 'Bearer ' + localStorage.getItem('token')
      },
      body: formData
    })

    if (response.ok) {
      ElMessage.success('商品添加成功')
      router.push('/admin')
    } else {
      const error = await response.json()
      ElMessage.error(error.message || '提交失败')
    }
  } catch (error) {
    ElMessage.error('网络请求失败')
  }
}

// 取消逻辑保持不变
const handleCancel = () => {
  router.push('/admin')
}
</script>

<style scoped>
.admin-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.form-wrapper {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

.form-container {
  width: 100%;
  max-width: 600px;
  background: #fff;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}
</style>