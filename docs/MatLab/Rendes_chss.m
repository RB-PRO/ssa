function Rendes_chss(RGB, Name)
    if length(RGB)<100
        return
    end
    try
        chss2(rgb2pw(RGB, "Cr"), "endh/"+Name+'/', Name);
    catch
        disp("ERROR: ������ ��� ������������ ������ �� �������� ������");
    end
    CloseFigure
end
