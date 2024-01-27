function [out, lower_prct, upper_prct] = RaznFilter(signal, threshold)
%     Nmed = 5;
%     signal = signal-medfilt1(signal, Nmed);
%     
%     % ������ �������� �������
%     len = length(signal);
%     FirstRaznSignal = [signal(1), signal];
%     FirstRaznSignal = FirstRaznSignal-[signal, signal(len)];
%     
%     disp("�������� ������� ��� insFrc_AcfNrm ��� ������� �����������: " + Nmed);  
%     prct = prctile(FirstRaznSignal, threshold);
% 
%     lower_prct=prct(1);
%     upper_prct=prct(2);
%     
%     disp("lower_prct " + lower_prct);
%     disp("upper_prct " + upper_prct);
%     
% %     signal=signal-medfilt1(signal, Nmed);
%     out = FirstRaznSignal;
%     
%     figure();
%     plot(FirstRaznSignal,'blue--'); hold on; grid on;
%     line('XData', [0 200], 'YData', [upper_prct upper_prct], 'Color','black','LineStyle','--');
%     line('XData', [0 200], 'YData', [lower_prct lower_prct], 'Color','black','LineStyle','--');
%     
%     index = 1;
%     MemoryValue = FirstRaznSignal(1);
%     
%     % ���� �� ����� �������
%     for value = FirstRaznSignal
%         if (value < lower_prct || value > upper_prct)
%             out(index) = MemoryValue;
%         end
%        MemoryValue=out(index);
%        index=index+1;
%     end
%     plot(out,'red'); grid on;
%     ylabel("Hz",'interp','none'); xlabel("ns",'interp','none'); title("������ ��������"); grid on;
%     legend('������ ��������','down','up','������������', 'Location', 'southoutside');
    

    % ���������� ������ ������� (������������� �����������)
    difference = diff(signal);

    % ����������� ������ ��� ����������� ��������
%     threshold = 5; % ������ ���������� �������� (�������� �� ���� �������� ��������)
    prct = prctile(difference, threshold);
    lower_prct=prct(1);
    upper_prct=prct(2);
    
    
%     figure('Position', [0 0 900 800]);
%     plot(difference,'blue'); hold on; grid on;
    len=length(signal);
%     line('XData', [0 len], 'YData', [upper_prct upper_prct], 'Color','black','LineStyle','--');
%     line('XData', [0 len], 'YData', [lower_prct lower_prct], 'Color','black','LineStyle','--');

%     line('XData', [0 200], 'YData', [threshold threshold], 'Color','green','LineStyle','--');
%     line('XData', [0 200], 'YData', [-threshold -threshold], 'Color','green','LineStyle','--');

%     % ���������� ����������� ����������� �� ������ ������
%     smoothedSignal = signal;
%     for i = 1:(length(difference))
%         xlim([i-5 i+5]);
%         disp(difference(j));
%         if i == 166
%             disp(166);
%         end
%         %if abs(difference(i)) > threshold
%         if (difference(i) < lower_prct || difference(i) > upper_prct)
%             % ���� ������� ��������� �����, ��������� �����������
% %             smoothedSignal(i+1) = (signal(i) + signal(i+1)) / 2; % ������� ����������� �� �������� ��������
%             j = i + 1;
%             while (difference(j) < lower_prct || difference(j) > upper_prct) && j ~= length(difference)-1
%                  j = j + 1;
%             end
%             smoothedSignal(i+1) = (signal(i) + signal(i+1)) / 2;
%         end
%     end

    %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
    % ���������� ������ ������� (������������� �����������)
% difference = diff(signal);

% ����������� ������ ��� ����������� ��������
threshold = 0.1; % ������ ���������� �������� (�������� �� ���� �������� ��������)

% ���������� ����������� ����������� �� ������ ������
smoothedSignal = signal;
for i = 1:length(difference)
    if abs(difference(i)) > threshold
        % ���� ������� ��������� �����, ��������� �����������
        smoothedSignal(i+1) = (signal(i) + signal(i+1)) / 2; % ������� ����������� �� �������� ��������
        % ��������� ��������� ������� �� ���������� ������ � ���������� ��
        j = i + 1;
        while j < length(difference) && abs(difference(j)) > threshold
            smoothedSignal(j+1) = (signal(j) + signal(j+1)) / 2; % ����������
            j = j + 1;
        end
    end
end
    %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
    
    out = smoothedSignal;
    out = difference;
    
%     plot(out,'red'); grid on;
%     plot(signal,'green--'); grid on;
%     ylabel("Hz",'interp','none'); xlabel("ns",'interp','none'); title("������ ��������"); grid on;
%     legend('������ ��������','down','up','������������', 'Location', 'southoutside');
    
    
end