function CloseFigure()
    closes=["Covariance matrix" "Eigenvalues" "Original time series and reconstruction" "��� ����������� ����� sET12 ��������� pw" "������������� ��� ����������� ����� sET12 ��������� pw" "������������� ������� sET12 ��������� pw" "Unwrape phase pulse wave" "Frequencie and energy pulse wave"]; 
    for icloses = 1:length(closes)
        try
            close(closes(icloses));
        catch
        end
    end
end
