import { GET } from '@/api/ajax';
import logo from '@/assets/img/logo.png';
import { SET_USERINFO } from '@/common/mutation-types';
import { USER_INFO } from '@/common/url';

// initial state
// shape: [{ id, quantity }]
const state = {
  userInfo: {
    openid: '',
    nickname: '',
    sex: '1',
    province: '',
    city: '',
    country: '',
    headimgurl: '',
  },
};

// getters
const getters = {};
// mutations
const mutations = {
  [SET_USERINFO](state, payload) {
    state.userInfo = payload;
  },
};
// actions
const actions = {
  initUserInfo({ commit }) {
    GET({
      url: USER_INFO,
      func: (response) => {
        const data = response.data;
        const userInfo = data.data;
        commit(SET_USERINFO, userInfo);
      },
      errFunc: (error) => {
        console.log(error);
        commit(SET_USERINFO, {
          openid: '',
          nickname: 'mu-mo',
          sex: '1',
          province: '',
          city: '',
          country: '',
          headimgurl: `${logo}`,
        });
      },
    });
  },
};


export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
