const registerForm = document.getElementById('register-form');

const registerUser = async (event) => {
  event.preventDefault();

  const name = document.getElementById('username').value;
  const password = document.getElementById('password').value;
  const email = document.getElementById('email').value;

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

  if (email) {
    reason = isValidPassword(email);
    if (reason) {
      alert('Login failed: ' + reason);
      return;
    }
    data.email = email;
  }

  try {
    const resp = await fetch(`/api/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (resp.ok) {
      alert('Registration successful!');
      window.location.href = '/login';
    } else {
      const errorText = await resp.text();
      alert('Registration failed: ' + errorText);
    }
  } catch (error) {
    console.error('Error:', error);
    alert('An error occurred during registration.');
  }
};

registerForm.addEventListener('submit', registerUser);
