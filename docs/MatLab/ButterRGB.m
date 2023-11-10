function RGB2 = ButterRGB(RGB, VideoFile) 

    % Задайте параметры фильтра 
    order = 4; % Порядок фильтра

    % Создайте фильтр Баттерворта
    [b, a] = butter(order, [0.2 0.9], 'stop');

    % Примените фильтр к RGB сигналу
    RGB2 = zeros(size(RGB));

    for channel = 1:3
        RGB2(:, channel) = filter(b, a, RGB(:, channel));
    end 

    SaveRGB(NameVideoFile(VideoFile)+'_but.txt', RGB2)
end