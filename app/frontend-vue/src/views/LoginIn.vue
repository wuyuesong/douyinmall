<template>
    <Header></Header>
    <div class="d-flex justify-content-center align-items-center" style="height: 50vh;">
      <div class="col-4">
        <form method="post" action="/auth/login?next={{ .next }}">
          <div class="mb-3">
            <label for="email" class="form-label">邮件</label>
            <input type="email" class="form-control" id="email" name="email">
          </div>
          <div class="mb-3">
            <label for="password" class="form-label">密码</label>
            <input type="password" class="form-control" id="password" name="password">
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

    export default {
    components: {
        Header
    }
    }
  </script>
  