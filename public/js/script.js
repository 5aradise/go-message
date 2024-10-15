const loginButton = document.getElementById('login-button');
const signupButton = document.getElementById('signup-button');
const usernameDisplay = document.getElementById('username-display');
const userBox = document.getElementById('user-box');
const signoutLink = document.getElementById('signout-link');
const chatSetup = document.getElementById('chat-setup');
const chatNameInput = document.getElementById('chat-name-input');
const createChatButton = document.getElementById('create-chat-button');
const connectChatButton = document.getElementById('connect-chat-button');

let clientName = '';

const authClient = async () => {
  const cookiesName = getCookie('name');
  if (!cookiesName) {
    if (await updateAuthorization()) {
      await authClient();
    }
    return;
  }
  clientName = cookiesName;
  usernameDisplay.textContent = clientName;
  loginButton.style.display = 'none';
  signupButton.style.display = 'none';
  userBox.style.display = 'flex';
  chatSetup.style.display = 'block';
};

const signout = async () => {
  try {
    const resp = await fetch(`/api/signout`, {
      method: 'POST',
    });

    if (!resp.ok) {
      if (await updateAuthorization()) {
        signout();
      } else {
        window.location.reload();
      }
      return;
    }
  } catch (error) {
    console.error('Error:', error);
    alert('An error occurred during signout.');
  }

  window.location.reload();
};

loginButton.onclick = () => {
  window.location.href = '/login';
};

signupButton.onclick = () => {
  window.location.href = '/signup';
};

signoutLink.onclick = signout;

createChatButton.onclick = async () => {
  const chatName = chatNameInput.value.trim();
  if (!chatName) {
    return;
  }

  try {
    const resp = await fetch(`/api/chats`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name: chatName }),
    });

    if (!resp.ok) {
      const errorText = await resp.text();
      alert('Creating chat failed: ' + errorText);
      return;
    }
  } catch (error) {
    console.error('Error:', error);
    alert('An error occurred during creating chat.');
  }

  window.location.href = `/chats/${chatName}`;
};

connectChatButton.onclick = () => {
  const chatName = chatNameInput.value.trim();
  if (chatName) {
    window.location.href = `/chats/${chatName}`;
  }
};

authClient();
