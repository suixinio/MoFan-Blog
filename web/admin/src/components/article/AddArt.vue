<template>
  <div>
    <a-card>
      <h3>{{ id ? '编辑文章' : '新增文章' }}</h3>
      <a-form-model :model='artInfo' ref='artInfoRef' :rules='artInfoRules' :hide-required-mark='true'>

        <a-form-model-item label='文章标题' prop='title'>
          <a-input style='width: 300px' v-model='artInfo.title'></a-input>
        </a-form-model-item>

        <a-form-model-item label='文章分类' prop='cid'>
          <a-select v-model='artInfo.cid' placeholder='请选择分类' style='width: 200px' @change='cateChange'>
            <a-select-option
              v-for='item in catelist'
              :key='item.ID'
              :value='item.ID'
            >{{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>

        <a-form-model-item label='文章描述' prop='desc'>
          <a-input type='textarea' v-model='artInfo.desc'></a-input>
        </a-form-model-item>

        <a-form-model-item label='文章缩略图' prop='img'>
          <a-upload
            name='file'
            :multiple='false'
            :action='upUrl'
            :headers='headers'
            @change='upChange'
            listType='picture'
            :defaultFileList='fileList'
          >
            <a-button>
              <a-icon type='upload' />
              Click to Upload
            </a-button>
            <br />
            <template v-if='id'>
              <img :src='artInfo.img' style='width: 120px;height: 100px'>
            </template>
          </a-upload>

        </a-form-model-item>

        <a-form-model-item label='文章内容' prop='content'>
          <Editor v-model='artInfo.content'></Editor>
        </a-form-model-item>

        <a-form-model-item>
          <a-button type='danger' style='margin-right: 15px;' @click='artOk(artInfo.id)'>
            {{ artInfo.id ? '更新' : '提交' }}
          </a-button>
          <a-button type='primary' @click='addCancel'>{{ artInfo.id ? '清空' : '取消' }}</a-button>
        </a-form-model-item>

      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugin/http'
import Editor from '../editor/index'

export default {
  components: {
    Editor
  },
  props: ['id'],
  data () {
    return {
      catelist: [],
      artInfo: {
        id: 0,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: ''
      },
      upUrl: Url + 'upload',
      headers: {},
      fileList: [],
      artInfoRules: {
        title: [{
          required: true,
          message: '请输入文章标题',
          trigger: 'blur'
        }],
        cid: [{
          required: true,
          message: '请输入文章分类',
          trigger: 'change'
        }],
        desc: [{
          required: true,
          message: '请输入文章描述',
          trigger: 'blur'
        }, {
          max: 120,
          message: '描述最多可写120个字符',
          trigger: 'blur'
        }],
        img: [{
          required: true,
          message: '请输入文章缩略图',
          trigger: 'blur'
        }],
        content: [{
          required: true,
          message: '请输入文章内容',
          trigger: 'blur'
        }]
      }
    }
  },
  created () {
    this.getCateList()
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem('token')}`
    }
    if (this.id) {
      this.getArtInfo(this.id)
    }
  },
  methods: {
    // 查询文章信息
    async getArtInfo (id) {
      const { data: res } = await this.$http.get(`article/${id}`)
      if (res.status !== 200) {
        if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
          window.sessionStorage.clear()
          await this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    },
    async getCateList () {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
          window.sessionStorage.clear()
          await this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.catelist = res.data
    },
    cateChange (value) {
      this.artInfo.cid = value
    },
    //  上传接口
    upChange (info) {
      if (info.file.status !== 'uploading') {
        console.log(info.file, info.fileList)
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} file uploaded successfully`)
        const imgUrl = info.file.response.url
        this.artInfo.img = imgUrl
        console.log(this.artInfo.img)
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} file upload failed.`)
      }
    },
    // 提交或者更新文章
    artOk (id) {
      this.$refs.artInfoRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        if (id === 0) {
          const { data: res } = await this.$http.post('article/add', this.artInfo)
          if (res.status !== 200) {
            if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
              window.sessionStorage.clear()
              await this.$router.push('/login')
            }
            this.$message.error(res.message)
          }
          this.$message.success('添加文章成功')
        } else {
          const { data: res } = await this.$http.put(`article/${id}`, this.artInfo)
          if (res.status !== 200) {
            if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
              window.sessionStorage.clear()
              await this.$router.push('/login')
            }
            this.$message.error(res.message)
          }
          this.$message.success('更新文章成功')
        }
      })
    },
    addCancel () {
      this.$refs.artInfoRef.resetFields()
      this.artInfo.id = 0
    }
  }
}
</script>

<style scoped>

</style>
