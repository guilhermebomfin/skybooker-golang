document.addEventListener("DOMContentLoaded", function() {
    const seatsContainer = document.getElementById('seatsContainer');

    // Obtenha o ID do voo da query string
    const urlParams = new URLSearchParams(window.location.search);
    const flightId = urlParams.get('voo');

    // Faça uma requisição à API para obter os assentos do voo específico
    fetch(`http://localhost:8080/flights/${flightId}/seats`)
        .then(response => response.json())
        .then(seats => {
            seats.forEach(seat => {
                const seatElement = document.createElement('div');
                seatElement.className = `seat ${seat.occupied ? 'occupied' : 'available'}`;
                seatElement.textContent = seat.number;

                // Adicione um evento de clique para selecionar o assento, se estiver disponível
                if (!seat.occupied) {
                    seatElement.addEventListener('click', function() {
                        // Lógica para selecionar o assento, se necessário
                        alert(`Assento ${seat.number} selecionado!`);
                    });
                }

                seatsContainer.appendChild(seatElement);
            });
        })
        .catch(error => console.error('Erro ao obter dados da API:', error));
});