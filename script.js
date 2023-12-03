document.addEventListener("DOMContentLoaded", function() {
    const tableBody = document.querySelector('#flightsTable tbody');
    const searchBar = document.getElementById('searchBar');

    fetch('http://localhost:8080/flights')
        .then(response => response.json())
        .then(data => {
            // Manter uma cópia dos dados originais para facilitar o filtro
            const originalData = [...data];

            // Preencher a tabela inicialmente com todos os voos
            fillTable(originalData);

            // Adicionar um ouvinte de evento à barra de pesquisa
            searchBar.addEventListener('input', function() {
                const searchTerm = searchBar.value.toLowerCase();
                const filteredData = filterFlights(originalData, searchTerm);
                fillTable(filteredData);
            });
        })
        .catch(error => console.error('Erro ao obter dados da API:', error));

    // Função para preencher a tabela com dados
    function fillTable(data) {
        // Limpar a tabela antes de preencher
        tableBody.innerHTML = '';

        data.forEach(flight => {
            const row = tableBody.insertRow();
            row.insertCell(0).textContent = flight.id;
            row.insertCell(1).textContent = flight.origin;
            row.insertCell(2).textContent = flight.destination;
            row.insertCell(3).textContent = flight.departure;

            // Adicionar botão para assentos
            const seatsButton = document.createElement('button');
            seatsButton.textContent = 'Ver Assentos';
            seatsButton.addEventListener('click', function() {
                // Redirecionar para a página de assentos do voo
                window.location.href = `seats.html?voo=${flight.id}`;
                displayReservationForm(seat, flight); // Passando o objeto do voo corretamente
            });
            
            

            // Adicionar status de disponibilidade dos assentos
            const availabilityStatus = document.createElement('span');
            availabilityStatus.textContent = ' (Disponíveis)'; // ou ' (Ocupados)' conforme necessário
            availabilityStatus.style.color = 'green'; // ou 'red' para indicar ocupados

            // Adicionar botão e status à célula
            const cell = row.insertCell(4);
            cell.appendChild(seatsButton);
            cell.appendChild(availabilityStatus);
        });
    }

    // Função para filtrar voos com base no termo de pesquisa
    function filterFlights(data, searchTerm) {
        return data.filter(flight => {
            const originMatch = flight.origin.toLowerCase().includes(searchTerm);
            const destinationMatch = flight.destination.toLowerCase().includes(searchTerm);
            const departureMatch = flight.departure.toLowerCase().includes(searchTerm);

            return originMatch || destinationMatch || departureMatch;
        });
    }
});
