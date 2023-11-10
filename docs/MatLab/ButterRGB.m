function RGB2 = ButterRGB(RGB, VideoFile) 

    % ������� ��������� ������� 
    order = 4; % ������� �������

    % �������� ������ �����������
    [b, a] = butter(order, [0.2 0.9], 'stop');

    % ��������� ������ � RGB �������
    RGB2 = zeros(size(RGB));

    for channel = 1:3
        RGB2(:, channel) = filter(b, a, RGB(:, channel));
    end 

    SaveRGB(NameVideoFile(VideoFile)+'_but.txt', RGB2)
end