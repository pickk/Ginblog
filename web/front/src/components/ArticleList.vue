<template>
  <!-- 主列，包含所有内容的栅格列 -->
  <v-col>
    <!-- 文章卡片，v-for 用于遍历 artList 数组，创建每篇文章的卡片 -->
    <!-- 遍历 artList 数组，生成每个文章的卡片 -->
    <v-card
        class="ma-3 d-flex flex-no-wrap justify-space-between align-center"
        v-for="item in artList"
        :key="item.id"
    link
    @click="goToDetail(item.ID)"
    >
    <!-- 每个卡片的唯一键，避免重复 -->
    <!-- 使卡片整个区域变成链接效果 -->
    <!-- 点击卡片时导航到文章详情页面 -->
    <!-- 头像组件，显示文章的图片（桌面视图时可见） -->
    <v-avatar class="ma-3 hidden-sm-and-down" size="125" tile>
      <!-- 图片组件，显示文章图片 -->
      <v-img :src="item.img"></v-img>
    </v-avatar>

    <!-- 内容列，包含文章的标题、描述、发布日期等 -->
    <v-col>
      <!-- 卡片标题，包含类别标签和文章标题 -->
      <v-card-title>
        <!-- 类别标签，显示文章类别 -->
        <v-chip color="pink" outlined label class="mr-3 white--text">
          {{ item.Category.name }} <!-- 显示文章类别名称 -->
        </v-chip>
        <!-- 文章标题 -->
        <div>{{ item.title }}</div>
      </v-card-title>

      <!-- 卡片副标题，显示文章的描述 -->
      <v-card-subtitle class="mt-1" v-text="item.desc"></v-card-subtitle>

      <!-- 分割线，用于分隔不同部分 -->
      <v-divider class="mx-4"></v-divider>

      <!-- 卡片文本，显示文章的元数据 -->
      <v-card-text class="d-flex align-center">
        <!-- 发布日期 -->
        <div class="d-flex align-center">
          <!-- 日历图标 -->
          <v-icon class="mr-1" small>{{ 'mdi-calendar-month' }}</v-icon>
          <!-- 格式化后的发布日期 -->
          <span>{{ item.CreatedAt | dateformat('YYYY-MM-DD') }}</span>
        </div>

        <!-- 评论数 -->
        <div class="mx-4 d-flex align-center">
          <!-- 评论图标 -->
          <v-icon class="mr-1" small>{{ 'mdi-comment' }}</v-icon>
          <!-- 评论数量 -->
          <span>{{ item.comment_count }}</span>
        </div>

        <!-- 阅读数 -->
        <div class="mx-1 d-flex align-center">
          <!-- 阅读图标 -->
          <v-icon class="mr-1" small>{{ 'mdi-eye' }}</v-icon>
          <!-- 阅读数量 -->
          <span>{{ item.read_count }}</span>
        </div>
      </v-card-text>
    </v-col>
    </v-card>

    <!-- 分页组件 -->
    <div class="text-center">
      <!-- 分页控件，显示页码 -->
      <v-pagination
          color="indigo"
      total-visible="7"
      v-model="queryParam.pagenum"
      :length="Math.ceil(total / queryParam.pagesize)"
      @input="getArtList()"
      ></v-pagination>
    </div>
  </v-col>
</template>
<script>

import { useRouter } from 'vue-router';

export default {
  // data 选项，用于定义组件的响应式数据
  setup() {
    const router = useRouter();

    // 导航函数
    const goToDetail = (id) => {
      router.push(`/article/detail/${id}`);
    };

    return { goToDetail };
  },
  data() {
    return {
      // 存储文章列表的数据
      artList: [],
      // 查询参数，用于分页请求
      queryParam: {
        pagesize: 5, // 每页显示的文章数量
        pagenum: 1   // 当前页码，默认为第 1 页
      },
      // 当前文章数（可以用来跟踪某些计数）
      count: 0,
      // 文章的总数（从服务器获取）
      total: 0
    }
  },

  // 生命周期钩子 created - 组件实例化时调用
  created() {
    // 此处未定义内容，通常可以在这里执行数据初始化或 API 调用
  },

  // 生命周期钩子 mounted - 组件挂载到 DOM 后调用
  mounted() {
    // 调用获取文章列表的方法
    this.getArtList()
  },

  // 过滤器选项（此处为空，若有需要可以定义过滤器）
  filters: {},

  // 方法选项，用于定义组件的函数和业务逻辑
  methods: {
    $router,
    // 获取文章列表的方法
    async getArtList() {
      // 发起 HTTP GET 请求获取文章数据
      const { data: res } = await this.$http.get('article', {
        params: {
          pagesize: this.queryParam.pagesize, // 每页的文章数
          pagenum: this.queryParam.pagenum    // 当前页码
        }
      })

      // 将获取到的数据赋值给 artList，用于渲染文章列表
      this.artList = res.data
      // 将获取到的文章总数赋值给 total，供分页使用
      this.total = res.total
    }
  }
}

</script>
<style lang=""></style>
