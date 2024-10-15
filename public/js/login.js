const registerForm = document.getElementById('login-form');

const loginUser = async (event) => {
  event.preventDefault();

  const name = document.getElementById('username').value;
  const password = document.getElementById('password').value;

  let reason;

  reason = isValidName(name);
  if (reason) {
    alert('Login failed: ' + reason);
    return;
  }
  reason = isValidPassword(password);
  if (reason) {
    alert('Login failed: ' + reason);
    return;
  }

  const data = {
    name,
    password,
  };

  try {
    const resp = await fetch(`/api/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (resp.ok) {
      alert('Login successful!');
      window.location.href = '/';
    } else {
      const errorText = await resp.text();
      alert('Login failed: ' + errorText);
    }
  } catch (error) {
    console.error('Error:', error);
    alert('An error occurred during login.');
  }
};

registerForm.addEventListener('submit', loginUser);
