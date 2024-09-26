const getCookie = (name) => {
  const cookies = document.cookie.split(';');

  for (const cookie of cookies) {
    const cookiePair = cookie.split('=');

    const cookieName = cookiePair[0].trim();

    if (cookieName === name) {
      return decodeURIComponent(cookiePair[1]);
    }
  }

  return null;
};

const minNameLen = 3;
const maxNameLen = 15;

const isValidName = (name) => {
  const l = name.length;
  if (l < minNameLen) {
    return `name must have at least ${minNameLen} characters`;
  }
  if (l > maxNameLen) {
    return `name must have no more than ${maxNameLen} characters`;
  }
  if (!/^[\p{L}\p{N}_-]+$/u.test(name)) {
    return "name must contains only letters, numbers and '_', '-'";
  }
  return null;
};

const minPasswordLen = 1;
const maxPasswordLen = 64;

const isValidPassword = (password) => {
  const l = password.length;
  if (l < minPasswordLen) {
    return `password must have at least ${minNameLen} characters`;
  }
  if (l > maxPasswordLen) {
    return `password must have no more than ${maxNameLen} characters`;
  }
  if (!/^[a-zA-Z0-9\p{P}\p{S}]+$/u.test(password)) {
    return 'password must contains only english letters numbers and special characters';
  }
  return null;
};

const isValidEmail = (email) => {
  if (!/^[\w.-]+@[a-zA-Z\d.-]+\.[a-zA-Z]{2,4}$/u.test(email)) {
    return 'invalide email';
  }
  return null;
};

const updateAuthorization = async () => {
  const username = getCookie('name');
  if (username) {
    return true;
  }
  try {
    const resp = await fetch(`http://${window.location.host}/api/refresh`, {
      method: 'POST',
    });
    if (resp.ok) {
      return true;
    } else {
      console.log(await resp.text());
    }
  } catch (error) {
    console.error('Error:', error);
    alert('An error occurred during updating access token.');
  }
  return false;
};
