document.addEventListener('DOMContentLoaded', () => {
    const loginBtn = document.getElementById('btn-login');
    const signupBtn = document.getElementById('btn-signup');
    const loginForm = document.getElementById('form-login');
    const signupForm = document.getElementById('form-signup');

    if (!loginBtn || !signupBtn || !loginForm || !signupForm) {
        console.error('Faltan botones o formularios');
        return;                 // ← aquí SÍ está permitido
    }

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