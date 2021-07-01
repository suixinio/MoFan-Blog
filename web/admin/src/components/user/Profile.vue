<template>
  <div>
    <a-card>
      <h3>个人设置</h3>
      <a-form-model :model='profileInfo' ref='profileInfoRef' :rules='profileInfoRules' :hide-required-mark='true'>

        <a-form-model-item label='作者名称' prop='name'>
          <a-input style='width: 300px' v-model='profileInfo.name'></a-input>
        </a-form-model-item>

        <a-form-model-item label='个人简介' prop='desc'>
          <a-input style='width: 300px' v-model='profileInfo.desc'></a-input>
        </a-form-model-item>
        <a-form-model-item label='个人简介背景图' prop='img'>
          <a-upload
            name='file'
            :multiple='false'
            :action='upUrl'
            :headers='headers'
            @change='upChangeImg'
            listType='picture'
            :defaultFileList='fileList'
          >
            <a-button>
              <a-icon type='upload' />
              Click to Upload
            </a-button>
            <br />
            <template v-if='profileInfo.img'>
              <img :src='profileInfo.img' style='width: 120px;height: 100px'>
            </template>
          </a-upload>

        </a-form-model-item>
        <a-form-model-item label='个人头像' prop='avatar'>
          <a-upload
            name='file'
            :multiple='false'
            :action='upUrl'
            :headers='headers'
            @change='upChangeAvatar'
            listType='picture'
            :defaultFileList='fileList'
          >
            <a-button>
              <a-icon type='upload' />
              Click to Upload
            </a-button>
            <br />
            <template v-if='profileInfo.avatar'>
              <img :src='profileInfo.avatar' style='width: 120px;height: 100px'>
            </template>
          </a-upload>

        </a-form-model-item>

        <a-form-model-item>
          <a-button type='danger' style='margin-right: 15px;' @click='upOk'>
            更新
          </a-button>
        </a-form-model-item>

      </a-form-model>

    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugin/http'

export default {
  name: 'Profile.vue',
  data () {
    return {
      upUrl: Url + 'upload',
      fileList: [],
      profileInfo: {
        id: 1,
        name: '',
        desc: '',
        img: '',
        avatar: ''
      },
      profileInfoRules: {
        name: [{
          required: true,
          message: '请输入用户名',
          trigger: 'blur'
        }],
        desc: [{
          required: true,
          message: '请输入用户描述',
          trigger: 'blur'
        }, {
          max: 120,
          message: '描述最多可写120个字符',
          trigger: 'blur'
        }],
        img: [{
          required: true,
          message: '请输入背景图片',
          trigger: 'blur'
        }],
        avatar: [{
          required: true,
          message: '请输入头像图片',
          trigger: 'blur'
        }]
      }
    }
  },
  created () {
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem('token')}`
    }
    this.getProfileInfo()
  },
  methods: {
    async getProfileInfo () {
      const { data: res } = await this.$http.get(`profile/${this.profileInfo.id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.profileInfo = res.data
      this.profileInfo.id = res.data.ID
    },
    //  上传接口
    upChangeImg (info) {
      if (info.file.status !== 'uploading') {
        console.log(info.file, info.fileList)
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} file uploaded successfully`)
        const imgUrl = info.file.response.url
        this.profileInfo.img = imgUrl
        console.log(this.profileInfo.img)
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} file upload failed.`)
      }
    },
    upChangeAvatar (info) {
      if (info.file.status !== 'uploading') {
        console.log(info.file, info.fileList)
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} file uploaded successfully`)
        const imgUrl = info.file.response.url
        this.profileInfo.avatar = imgUrl
        console.log(this.profileInfo.img)
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} file upload failed.`)
      }
    },
    upOk () {
      this.$refs.profileInfoRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        console.log(this.profileInfo.id)
        const { data: res } = await this.$http.put(`profile/${this.profileInfo.id}`, this.profileInfo)
        if (res.status !== 200) {
          if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
            window.sessionStorage.clear()
            await this.$router.push('/login')
          }
          this.$message.error(res.message)
        }
        this.$message.success('更新用户信息成功')
      })
    }
  }
}
</script>

<style scoped>

</style>
