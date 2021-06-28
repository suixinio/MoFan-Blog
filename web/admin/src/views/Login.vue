<template>
  <div class='container'>
    <div class='loginBox'>
      登录页面
      <a-form-model class='loginForm'>
        <a-form-model-item>
          <a-input></a-input>
        </a-form-model-item>

        <a-form-model-item>
          <a-input></a-input>
        </a-form-model-item>

        <a-form-model-item class='loginBtn'>
          <a-button type='primary'>登录</a-button>
          <a-button type='info'>取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>

  </div>
</template>

<script>
export default {
  name: 'Login'
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}

</style>
