// 过期时间，默认7天
const age = 7 * 24 * 60 * 60 * 1000;
const getValue = (key) => {
  const jsonValue = localStorage.getItem(key);
  if (jsonValue) {
    return JSON.parse(jsonValue);
  }
  return null;
};
const isExpire = (key) => {
  const value = getValue(key);
  if (value) {
    const now = new Date().getTime();
    if (value.expireTime > now) {
      return false;
    }
    localStorage.removeItem(key);
  }
  return true;
};
const set = (key, value, expire) => {
  localStorage.removeItem(key);
  const now = new Date().getTime();
  let expireNum = age;
  if (expire !== undefined) {
    expireNum = parseInt(expire, 10);
    if (isNaN(expireNum)) {
      return;
    }
  }

  const newValue = {};
  newValue.value = value;
  // 加入时间
  newValue.createTime = now;
  // 过期时间
  newValue.expireTime = now + expireNum;
  localStorage.setItem(key, JSON.stringify(newValue));
};

const get = (key) => {
  if (!isExpire(key)) {
    const value = getValue(key);
    return value.value;
  }
  return null;
};


export {
  get,
  set,
  isExpire,
};
