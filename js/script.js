// Detectar si el dispositivo es móvil
function isMobileDevice() {
    return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent);
}

// Mostrar un mensaje de error si el dispositivo es móvil
if (isMobileDevice()) {
    alert("Lo sentimos, este código no está diseñado para funcionar en dispositivos móviles. Por favor, utiliza un dispositivo de escritorio.");
    document.body.innerHTML = `
        <div id="error-message" style="text-align: center; margin-top: 20vh; font-family: Arial, sans-serif;">
            <h1>¡Lo sentimos!</h1>
            <p>Esta aplicación no es compatible con dispositivos móviles.</p>
            <p>Por favor, accede desde un dispositivo de escritorio.</p>
        </div>
    `;
} else {
    // Asignar funcionalidad al botón de descarga
    document.getElementById("download-button").onclick = function () {
        window.location.href = "/download-json";
    };
}
