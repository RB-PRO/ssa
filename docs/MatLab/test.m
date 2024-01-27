% �������� ��������� �������
signal = randn(1, 1000); % ��������� ���������� ������� �� 1000 �����

% ����������� ����������� �� 1 �� 99
percentiles = prctile(signal, 1:99);

% ���������� ����������
filtered_signal = signal(signal >= percentiles(25) & signal <= percentiles(75));

% ����������� �����������
figure;
subplot(2,1,1);
plot(signal, 'b'); % �������� ������
title('�������� ������');

subplot(2,1,2);
plot(filtered_signal, 'r'); % ��������������� ������
title('��������������� ������');
