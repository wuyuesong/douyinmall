<template>
    <Header></Header>
    <div class="d-flex justify-content-center align-items-center" style="height: 50vh;">
      <div class="col-4">
        <form @submit.prevent="handleLogin">
          <div class="mb-3">
            <label for="email" class="form-label">邮件</label>
            <input type="email" class="form-control" id="email" name="email" v-model="email"  >
          </div>
          <div class="mb-3">
            <label for="password" class="form-label">密码</label>
            <input type="password" class="form-control" id="password" name="password" v-model="password">
          </div>
          <div class="mb-3">
            如果没有账户, 点击这里 <a href="/sign-up">注册</a>.
          </div>
          <button type="submit" class="btn btn-primary">登录</button>
        </form>
      </div>
    </div>
  </template>
  
<script>
    import '../assets/bootstrap.min.css';
    import Header from '../components/Header.vue';
    import { useRouter } from 'vue-router'
    import Cookies from 'js-cookie'

    const router = useRouter()
    export default {
        data() {
            return {
                email: '',
                password: '',
                redirect: null
            }
        },
        components: {
            Header
        },
        created() {
            // 获取路由中的 redirect 参数
            this.redirect = this.$route.query.redirect
            console.log('this.redirect', this.redirect)
        },
        methods: {
            async handleLogin() {
                try {
                    // 这里替换为实际的登录请求
                    const response = await this.$axios.post('/auth/login', {
                        email: this.email,
                        password: this.password
                    })

                    if (response.code == 200) { 
                      this.$message({
                          message: '登录成功！',
                          type: 'success',
                          duration: 1500
                      })
                      localStorage.setItem('token', response.token); // 写入 localStorage
                      localStorage.setItem('cartNum', response.cartNum.toString())
                      this.$store.commit('SET_CART_NUM', response.cartNum.toString())
                      // 登录成功后跳转
                      console.log('this.redirect', this.redirect)
                      this.$router.push(this.redirect || '/')
                    } else {
                      this.$message({
                          message: '登录失败！',
                          type: 'error',
                          duration: 1500
                      })
                    }

                } catch (error) {
                    console.error('登录失败', error)
                }
            }
        }
    }
</script>
  