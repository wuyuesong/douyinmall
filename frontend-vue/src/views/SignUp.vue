<template>
    <Header></Header>
    <div class="d-flex justify-content-center align-items-center" style="height: 50vh;">
        <div class="col-4">
          <form @submit.prevent="handleSignUp">
            <div class="mb-3">
              <label for="email" class="form-label">邮件 <Required /></label>
              <input type="email" class="form-control" id="email" v-model="email" name="email">
            </div>
            <div class="mb-3">
              <label for="password" class="form-label">密码 <Required /></label>
              <input type="password" class="form-control" id="password" v-model="password" name="password">
            </div>
            <div class="mb-3">
                <label for="password_confirm" class="form-label">密码确认 <Required /></label>
                <input type="password" class="form-control" id="password_confirm" v-model="password_confirm" name="password_confirm">
              </div>
            <div class="mb-3">
                已有账号，点击此处登录 <a href="/login">登录</a>.
            </div>
            <button type="submit" class="btn btn-primary">注册</button>
          </form>
    </div>
</div>
</template>

<script>
    import '../assets/bootstrap.min.css';
    import Header from '../components/Header.vue';
    import Required from '../components/Required.vue';

    export default {
        components: {
            Header,
            Required
        },
        data() {
            return {
                email: '',
                password: '',
                password_confirm: ''
            }
        },
        methods: {
            async handleSignUp() {
                try {
                    const response = await this.$axios.post('/auth/register', {
                        email: this.email,
                        password: this.password,
                        password_confirm: this.password_confirm
                    });
                    this.$router.push('/login');
                } catch (error) {
                    console.error('注册失败', error);
                }
            }
        }
    }
</script>