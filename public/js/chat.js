const messageInput = document.getElementById('message-input');
const sendButton = document.getElementById('send-button');
const messageArea = document.getElementById('message-area');
const usernameDisplay = document.getElementById('username-display');
const userBox = document.getElementById('user-box');
const chatNameHeader = document.getElementById('chat-name');

let ws = null;
let clientName = '';

const initializeWebSocket = async (chatName) => {
  ws = new WebSocket(`/api/ws/${chatName}`);
  ws.onerror = async () => {
    if (!clientName) {
      if (await updateAuthorization()) {
        initializeWebSocket(chatName);
      } else {
        alert('Please log in.');
      }
    } else {
      window.location.href = '/';
    }
  };
  ws.onmessage = (event) => {
    const message = JSON.parse(event.data);
    if (message instanceof Array) {
      for (const msg of message) {
        addMessage(
          msg.sender,
          msg.body,
          new Date(msg.time).toLocaleTimeString()
        );
      }
      return;
    }
    addMessage(
      message.sender,
      message.body,
      new Date(message.time).toLocaleTimeString()
    );
  };
};

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
  userBox.style.display = 'flex';
  messageInput.disabled = false;
  sendButton.disabled = false;
  const chatName = getChatNameFromURL();
  await initializeWebSocket(chatName);
  chatNameHeader.textContent = chatName;
  document.title = `Chat Room - ${chatName}`;
};

const getChatNameFromURL = () => {
  const pathArray = window.location.pathname.split('/');
  return pathArray[pathArray.length - 1];
};

const addMessage = (sender, body, time) => {
  const messageDiv = document.createElement('div');
  messageDiv.className = 'message' + (sender === clientName ? ' client' : '');

  const senderDiv = document.createElement('div');
  senderDiv.className = 'sender';
  senderDiv.textContent = sender;

  const textDiv = document.createElement('div');
  textDiv.className = 'text';
  textDiv.textContent = body;

  const timeDiv = document.createElement('div');
  timeDiv.className = 'time';
  timeDiv.textContent = time;

  messageDiv.appendChild(senderDiv);
  messageDiv.appendChild(textDiv);
  messageDiv.appendChild(timeDiv);

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
};

sendButton.onclick = sendMessage;

messageInput.addEventListener('keydown', (event) => {
  if (event.key === 'Enter') {
    sendMessage();
    event.preventDefault();
  }
});

authClient();
