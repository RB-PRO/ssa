% Создание тестового сигнала
signal = randn(1, 1000); % Генерация случайного сигнала из 1000 точек

% Определение процентилей от 1 до 99
percentiles = prctile(signal, 1:99);

% Применение фильтрации
filtered_signal = signal(signal >= percentiles(25) & signal <= percentiles(75));

% Отображение результатов
figure;
subplot(2,1,1);
plot(signal, 'b'); % Исходный сигнал
title('Исходный сигнал');

subplot(2,1,2);
plot(filtered_signal, 'r'); % Отфильтрованный сигнал
title('Отфильтрованный сигнал');
