const messageInput = document.getElementById('message-input');
const sendButton = document.getElementById('send-button');
const messageArea = document.getElementById('message-area');
const loginButton = document.getElementById('login-button');
const usernameDisplay = document.getElementById('username-display');
const unloggedMessage = document.getElementById('unlogged-message');
const userBox = document.getElementById('user-box');
const signoutLink = document.getElementById('signout-link');

const ws = new WebSocket(`ws://${window.location.host}/api/ws`);
ws.binaryType = 'arraybuffer';

const textDecoder = new TextDecoder();
const senderMsgDiv = 0;

let clientName = '';

const authClient = async () => {
  const cookiesName = getCookie('name');
  if (!cookiesName) {
    unloggedMessage.style.display = 'inline';
    return;
  }

  clientName = cookiesName;
  usernameDisplay.textContent = clientName;
  loginButton.style.display = 'none';
  userBox.style.display = 'flex';
  messageInput.disabled = false;
  sendButton.disabled = false;
};

const addMessage = (sender, body, isClient) => {
  const messageDiv = document.createElement('div');
  messageDiv.className = 'message' + (isClient ? ' client' : '');

  const senderDiv = document.createElement('div');
  senderDiv.className = 'sender';
  senderDiv.textContent = sender;

  const textDiv = document.createElement('div');
  textDiv.className = 'text';
  textDiv.textContent = body;

  messageDiv.appendChild(senderDiv);
  messageDiv.appendChild(textDiv);

  messageArea.appendChild(messageDiv);

  messageArea.scrollTop = messageArea.scrollHeight;
};

const sendMessage = () => {
  const msgText = messageInput.value.trim();
  if (!msgText) {
    return;
  }
  messageInput.value = '';

  ws.send(msgText);

  addMessage(clientName, msgText, true);
};

ws.onmessage = (event) => {
  const data = new Uint8Array(event.data);
  const divIndex = data.indexOf(senderMsgDiv);
  const sender = textDecoder.decode(data.slice(0, divIndex));
  const body = textDecoder.decode(data.slice(divIndex + 1));
  addMessage(sender, body, false);
};

loginButton.onclick = () => {
  window.location.href = '/login';
};

signoutLink.onclick = async () => {
  try {
    const resp = await fetch(`http://${window.location.host}/api/signout`, {
      method: 'POST',
    });

    if (!resp.ok) {
      const errorText = await resp.text();
      alert('Signout failed: ' + errorText);
      return;
    }
  } catch (error) {
    console.error('Error:', error);
    alert('An error occurred during login.');
  }

  window.location.reload();
};

sendButton.onclick = sendMessage;

messageInput.addEventListener('keydown', (event) => {
  if (event.key === 'Enter') {
    sendMessage();
    event.preventDefault();
  }
});

window.onload = authClient