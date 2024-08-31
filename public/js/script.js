const messageInput = document.getElementById('message-input');
const sendButton = document.getElementById('send-button');
const messageArea = document.getElementById('message-area');

const ws = new WebSocket(`ws://${window.location.host}/api/ws`);
ws.binaryType = 'arraybuffer';

const textDecoder = new TextDecoder();
const senderMsgDiv = 0;

let clientName = '';
let accessKey = '';

const authClient = async () => {
  while (true) {
    let newName = prompt('Enter name');
    if (!newName) {
      alert('Empty name!');
      continue;
    }

    newName = newName.trim();
    if (!newName) {
      alert('Empty name!');
      continue;
    }

    const resp = await fetch(`http://${window.location.host}/api/register`, {
      method: 'POST',
      body: JSON.stringify({ name: newName }),
    });
    if (resp.status !== 201) {
      alert(await resp.text());
      continue;
    }

    const jsonKey = await resp.json();
    accessKey = jsonKey.key;
    clientName = newName;
    break;
  }
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

ws.onopen = async () => {
  await authClient();
  ws.send(accessKey);
};

ws.onmessage = (event) => {
  const data = new Uint8Array(event.data);
  const divIndex = data.indexOf(senderMsgDiv);
  const sender = textDecoder.decode(data.slice(0, divIndex));
  const body = textDecoder.decode(data.slice(divIndex + 1));
  addMessage(sender, body, false);
};

sendButton.onclick = sendMessage;

messageInput.addEventListener('keydown', function (event) {
  if (event.key === 'Enter') {
    sendMessage();
    event.preventDefault();
  }
});
