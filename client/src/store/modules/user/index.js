import { GET } from '@/api/ajax';
import logo from '@/assets/img/logo.png';
import { SET_USERINFO } from '@/common/mutation-types';
import { USER_INFO } from '@/common/url';
import { set as setLocalStorage, get as getLocalStorage } from '@/util/localStorage';

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
    setLocalStorage('userInfo', payload);
    setLocalStorage('status', '3');
  },
};
// actions
const actions = {
  initUserInfo({ commit }) {
    let userInfo = getLocalStorage('userInfo');
    if (!userInfo) {
      GET({
        url: USER_INFO,
        func: (response) => {
          const data = response.data;
          userInfo = data.data;
        },
        errFunc: (error) => {
          console.log(error);
          userInfo = {
            openid: '',
            nickname: 'mu-mo',
            sex: '1',
            province: '',
            city: '',
            country: '',
            headimgurl: `${logo}`,
          };
        },
      }).then(() => {
        commit(SET_USERINFO, userInfo);
      });
    } else {
      commit(SET_USERINFO, userInfo);
    }
  },
};


export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
