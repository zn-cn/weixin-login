<template>
  <div>
    <div>
      <img :src="userInfo.headimgurl" />
    </div>
    <div>
      <p>{{userInfo.nickname}}</p>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  computed: {
    ...mapState('user', ['userInfo']),
  },
  methods: {
    ...mapActions('user', ['initUserInfo']),
  },
  created() {
    this.initUserInfo();
  },
};
</script>

<style>
</style>
