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

        <el-form-item label="商品描述" prop="description">
          <el-input v-model="form.description" type="textarea" />
        </el-form-item>

        <el-form-item label="商品分类" prop="categories">
          <el-select v-model="form.categorie" placeholder="请选择分类" style="width: 100%">
            <el-option label="普通商品" value="normal" />
            <el-option label="折扣商品" value="discount" />
          </el-select>
        </el-form-item>

        <el-form-item label="商品图片">
          <el-upload
            class="upload-demo"
            :action="uploadUrl"
            :headers="uploadHeaders"
            :on-success="handleImageSuccess"
            :file-list="fileList"
            list-type="picture"
            accept="image/*"
            :limit="1"
          >
            <el-button type="primary">点击上传</el-button>
            <template #tip>
              <div class="el-upload__tip">
                请上传JPEG/PNG格式的图片，大小不超过2MB（仅支持单张图片）
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
import type { UploadFile } from 'element-plus'
import { useRouter } from 'vue-router'
import type { FormInstance } from 'element-plus'

const router = useRouter()
const uploadUrl = import.meta.env.VITE_API_BASE + '/upload-images' // 替换为实际图片上传接口
const uploadHeaders = {
  Authorization: 'Bearer ' + localStorage.getItem('token')
}

// 表单数据结构
const form = reactive({
  categorie: 'normal',
  name: '',
  price: '',
  stock: '',
  description: '',
  picture: ''
})

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
  description: [
    { required: true, message: '请输入商品描述', trigger: 'blur' },
    { min: 10, max: 500, message: '长度在10到500个字符', trigger: 'blur' }
  ],
  categorie: [
    { required: true, message: '请选择商品分类', trigger: 'change' }
  ],
})

const handleImageSuccess = (response: any) => {
  if (response && response.url) {
    form.picture = response.url
    ElMessage.success('图片上传成功')
  }
}

const onSubmit = async () => {
  if (!form.picture) {
    ElMessage.error('请上传商品图片')
    return
  }

  const productData = {
    name: form.name,
    price: parseFloat(form.price),
    stock: parseInt(form.stock, 10),
    description: form.description,
    categorie: form.categorie,
    picture: form.picture
  }

  try {
    const response = await fetch(import.meta.env.VITE_API_BASE + '/product', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('token')
      },
      body: JSON.stringify(productData)
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