<template>
    <el-dialog
      v-model="dialogVisible"
      :title="formData.id ? '编辑地址' : '新增地址'"
      width="600px"
      @closed="handleClose"
    >
      <el-form 
        ref="formRef" 
        :model="formData" 
        :rules="rules" 
        label-width="80px"
        label-position="left"
      >
        <!-- 收货人 -->
        <el-form-item label="收货人" prop="consignee">
          <el-input 
            v-model="formData.consignee" 
            placeholder="请输入收货人姓名"
            clearable
          />
        </el-form-item>
  
        <!-- 手机号 -->
        <el-form-item label="手机号" prop="phone">
          <el-input
            v-model="formData.phone"
            placeholder="请输入11位手机号码"
            clearable
            maxlength="11"
          />
        </el-form-item>
  
        <!-- 省市区选择 -->
        <el-form-item label="所在地区" prop="region">
          <el-cascader
            v-model="formData.region"
            :options="regionOptions"
            :props="regionProps"
            placeholder="请选择省/市/区"
            style="width: 100%"
            clearable
          />
        </el-form-item>
  
        <!-- 详细地址 -->
        <el-form-item label="详细地址" prop="detail">
          <el-input
            v-model="formData.detail"
            placeholder="街道门牌、楼层房间号等信息"
            clearable
          />
        </el-form-item>
  
        <!-- 设为默认 -->
        <el-form-item label="设为默认" prop="isDefault">
          <el-switch v-model="formData.isDefault" />
        </el-form-item>
      </el-form>
  
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button 
            type="primary" 
            @click="handleSubmit"
            :loading="submitting"
          >
            保存地址
          </el-button>
        </div>
      </template>
    </el-dialog>
  </template>
  
  <script setup>
  import { ref, watch, reactive } from 'vue'
  import { ElMessage } from 'element-plus'
  import { regionData } from 'element-china-area-data'
  
  const props = defineProps({
    modelValue: Boolean,
    editData: {
      type: Object,
      default: null
    }
  })
  
  const emit = defineEmits(['update:modelValue', 'success'])
  
  // 弹窗显示控制
  const dialogVisible = ref(false)
  watch(
    () => props.modelValue,
    (val) => (dialogVisible.value = val)
  )
  watch(dialogVisible, (val) => emit('update:modelValue', val))
  
  // 表单数据
  const formData = reactive({
    consignee: '',
    phone: '',
    region: [],
    detail: '',
    isDefault: false
  })
  
  // 表单引用
  const formRef = ref(null)
  
  // 地区选择配置
  const regionOptions = regionData
  const regionProps = {
    label: 'label',
    value: 'value',
    children: 'children'
  }
  
  // 验证规则
  const rules = {
    consignee: [
      { required: true, message: '收货人不能为空', trigger: 'blur' },
      { min: 2, max: 10, message: '姓名长度2-10个字符', trigger: 'blur' }
    ],
    phone: [
      { required: true, message: '手机号不能为空', trigger: 'blur' },
      { 
        pattern: /^1[3-9]\d{9}$/,
        message: '请输入正确的手机号码',
        trigger: 'blur'
      }
    ],
    region: [
      { 
        required: true,
        validator: (_, value, callback) => {
          value?.length === 3 
            ? callback() 
            : callback(new Error('请选择完整的省市区'))
        },
        trigger: 'change'
      }
    ],
    detail: [
      { required: true, message: '详细地址不能为空', trigger: 'blur' },
      { min: 5, max: 50, message: '长度5-50个字符', trigger: 'blur' }
    ]
  }
  
  // 提交状态
  const submitting = ref(false)
  
  // 监听编辑数据变化
  watch(() => props.editData, (val) => {
    if (val) {
      Object.assign(formData, {
        ...val,
        region: val.province ? [val.province, val.city, val.district] : []
      })
    }
  })
  
  // 提交处理
  const handleSubmit = async () => {
    try {
      const valid = await formRef.value.validate()
      if (!valid) return
  
      submitting.value = true
      
      // 构造请求数据
      const requestData = {
        ...formData,
        province: formData.region[0],
        city: formData.region[1],
        district: formData.region[2]
      }
      
      // 调用API（根据实际接口调整）
      const api = formData.id 
        ? $axios.put(`/address/${formData.id}`, requestData)
        : $axios.post('/address', requestData)
      
      await api
      ElMessage.success('地址保存成功')
      emit('success')
      dialogVisible.value = false
    } catch (error) {
      if (error?.response?.data) {
        ElMessage.error(error.response.data.message)
      }
    } finally {
      submitting.value = false
    }
  }
  
  // 弹窗关闭时重置表单
  const handleClose = () => {
    formRef.value?.resetFields()
    Object.assign(formData, {
      consignee: '',
      phone: '',
      region: [],
      detail: '',
      isDefault: false
    })
  }
  </script>
  
  <style scoped>
  .dialog-footer {
    text-align: right;
  }
  </style>