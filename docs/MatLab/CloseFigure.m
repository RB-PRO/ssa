function CloseFigure()
    closes=["Covariance matrix" "Eigenvalues" "Original time series and reconstruction" "АКФ сингулярных троек sET12 сегментов pw" "Нормированные АКФ сингулярных троек sET12 сегментов pw" "Периодограмма Томсона sET12 сегментов pw" "Unwrape phase pulse wave" "Frequencie and energy pulse wave"]; 
    for icloses = 1:length(closes)
        try
            close(closes(icloses));
        catch
        end
    end
end
