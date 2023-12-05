document.addEventListener("DOMContentLoaded", function () {
    const seatsContainer = document.getElementById('seatsContainer');
    const reservationForm = document.getElementById('reservationForm');

    // Obtenha o ID do voo da query string
    const urlParams = new URLSearchParams(window.location.search);
    const flightId = urlParams.get('voo');

    // Variáveis para armazenar informações do voo
    let flightOrigin = '';
    let flightDestination = '';

    // Faça uma requisição à API para obter os assentos do voo específico
    fetch(`http://localhost:8080/flights/${flightId}/seats`)
        .then(response => response.json())
        .then(seats => {
            seats.forEach(seat => {
                const seatElement = document.createElement('div');
                seatElement.className = `seat ${seat.occupied ? 'occupied' : 'available'}`;
                seatElement.textContent = seat.number;

                // Adicione um evento de clique para exibir o formulário ao selecionar o assento, se estiver disponível
                if (!seat.occupied) {
                    seatElement.addEventListener('click', function () {
                        displayReservationForm(seat);
                    });
                }

                seatsContainer.appendChild(seatElement);
            });
        })
        .catch(error => console.error('Erro ao obter dados da API:', error));

    // Adicione um evento de submit ao formulário de reserva
    reservationForm.addEventListener('submit', function (event) {
        event.preventDefault();

        // Obtenha os dados do formulário
        const seatNumber = document.getElementById('seatNumber').value;
        const passengerName = document.getElementById('passengerName').value;
        const passengerDOB = document.getElementById('passengerDOB').value;
        const seatPrice = document.getElementById('seatPrice').value;

        // Atualize o assento com os dados do passageiro
        updateSeatWithReservation(seatNumber, passengerName, passengerDOB, seatPrice);

        // Oculte o formulário após a reserva
        reservationForm.style.display = 'none';

        // Recarregue a página automaticamente
        window.location.reload();
    });

    // Função para exibir o formulário de reserva
    function displayReservationForm(seat) {
        reservationForm.style.display = 'block';

        // Limpar campos existentes
        while (reservationForm.firstChild) {
            reservationForm.removeChild(reservationForm.firstChild);
        }

        // Adicionar o número do assento ao formulário
        const seatNumberField = document.createElement('input');
        seatNumberField.type = 'hidden';
        seatNumberField.name = 'seatNumber';
        seatNumberField.id = 'seatNumber';
        seatNumberField.value = seat.number;

        // Adicionar o preço ao formulário
        const priceLabel = document.createElement('label');
        priceLabel.textContent = 'Preço:';
        const priceInput = document.createElement('input');
        priceInput.type = 'text';
        priceInput.name = 'seatPrice';
        priceInput.id = 'seatPrice';
        priceInput.value = seat.price.toFixed(2);
        priceInput.readOnly = true;

        // Adicionar campos do formulário
        const nameLabel = document.createElement('label');
        nameLabel.textContent = 'Nome do Passageiro:';
        const nameInput = document.createElement('input');
        nameInput.type = 'text';
        nameInput.name = 'passengerName';
        nameInput.id = 'passengerName';
        nameInput.required = true;

        const dobLabel = document.createElement('label');
        dobLabel.textContent = '\nData de Nascimento:';
        const dobInput = document.createElement('input');
        dobInput.type = 'date';
        dobInput.name = 'passengerDOB';
        dobInput.id = 'passengerDOB';
        dobInput.required = true;

        const submitButton = document.createElement('button');
        submitButton.type = 'submit';
        submitButton.textContent = 'Finalizar Compra';

        // Adicionar campos ao formulário
        reservationForm.appendChild(seatNumberField);
        reservationForm.appendChild(priceLabel);
        reservationForm.appendChild(priceInput);
        reservationForm.appendChild(nameLabel);
        reservationForm.appendChild(nameInput);
        reservationForm.appendChild(dobLabel);
        reservationForm.appendChild(dobInput);
        reservationForm.appendChild(submitButton);
    }

    // Função para atualizar o assento com os dados da reserva
    function updateSeatWithReservation(seatNumber, passengerName, passengerDOB, seatPrice) {
        // Lógica para atualizar o assento na API com os dados fornecidos
        fetch(`http://localhost:8080/flights/${flightId}/seats/${seatNumber}`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                occupied: true,
                passengerName: passengerName,
                passengerDOB: passengerDOB,
                price: parseFloat(seatPrice), // Converte o preço de string para float
            }),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`Erro HTTP! Status: ${response.status}`);
                }
                // Trate a resposta como JSON
                return response.json();
            })
            .then(updatedSeat => {
                console.log(`Assento ${updatedSeat.number} reservado para ${updatedSeat.passengerName}`);
                
            })
            .catch(error => console.error('Erro ao atualizar o assento na API:', error));
    }
});
