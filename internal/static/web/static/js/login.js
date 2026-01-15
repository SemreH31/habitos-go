document.addEventListener('DOMContentLoaded', () => {
    const loginBtn = document.getElementById('btn-login');
    const signupBtn = document.getElementById('btn-signup');
    const loginForm = document.getElementById('form-login');
    const signupForm = document.getElementById('form-signup');

    if (!loginBtn || !signupBtn || !loginForm || !signupForm) {
        console.error('Faltan botones o formularios');
        return;                 // ← aquí SÍ está permitido
    }
    loginBtn.classList.add('active');
    loginForm.classList.add('active');
    signupForm.classList.remove('active');

    loginBtn.addEventListener('click', () => {
        loginForm.classList.add('active');
        signupForm.classList.remove('active');
        loginBtn.classList.add('active');
        signupBtn.classList.remove('active');
    });

    signupBtn.addEventListener('click', () => {
        signupForm.classList.add('active');
        loginForm.classList.remove('active');
        signupBtn.classList.add('active');
        loginBtn.classList.remove('active');
    });

    /* ----------  toggle password  ---------- */
    function togglePwd(button) {
        const input = button.previousElementSibling;
        const icon = button.querySelector('span');
        if (input.type === 'password') {
            input.type = 'text';
            icon.textContent = 'visibility';
        } else {
            input.type = 'password';
            icon.textContent = 'visibility_off';
        }
    }

    document.querySelectorAll('.toggle-pwd').forEach(btn =>
        btn.addEventListener('click', () => togglePwd(btn))
    );
});
document.getElementById('form-signup').addEventListener('submit', async (e) => {
    e.preventDefault();
    const data = Object.fromEntries(new FormData(e.target));
    if (data.password !== data.password_confirm) {
        alert("Passwords do not match");
        return;
    }

    const res = await fetch('/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    });

    const out = await res.json();
    if (res.ok) {
        alert(out.message);
        e.target.reset();
        document.getElementById('btn-login').click(); // ve a login
    } else {
        alert(out.error);
    }
});