const pathSegments = window.location.pathname.split('/');
const uniqueCode = pathSegments[pathSegments.length - 1];

document.getElementById('resetPasswordForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const password = document.getElementById('password').value;
    const password_confirmation = document.getElementById('password_confirmation').value;

    try {
        const response = await fetch(`/reset-password/${uniqueCode}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ password, password_confirmation }),
        });

        const result = await response.json();

        if (response.ok) {
            document.getElementById('message').textContent = 'Password reset successfully!';
        } else {
            document.getElementById('message').textContent = result.message || 'Failed to reset password.';
            document.getElementById('message').classList.add('danger');
        }
    } catch (error) {
        document.getElementById('message').textContent = 'An error occurred. Please try again.';
    }
});