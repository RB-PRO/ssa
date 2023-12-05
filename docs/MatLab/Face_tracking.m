function RGB = Face_tracking(VideoFile)
    vid_obj = VideoReader(strrep(VideoFile,"_nc",""));
    VideoFile = NameVideoFile(VideoFile);
    
    %



    FDetect = vision.CascadeObjectDetector('FrontalFaceCART','MergeThreshold',10); %Viola-Jones Algorithm
    %EyeDetect = vision.CascadeObjectDetector('EyePairSmall','UseROI', true);

    videoFrame = readFrame(vid_obj);
    bbox = step(FDetect, videoFrame);

    videoFrame = insertShape(videoFrame, "rectangle", bbox);
%     figure(); imshow(videoFrame); title("Detected face");

    %РџРѕРїС‹С‚РєРё СЃРµРіРјРµРЅС‚Р°С†РёРё РёР·РѕР±СЂР°Р¶РµРЅРёР№
    %РњРµС‚РѕРґ 1
    % Face_cluster = mean_shift(videoFrame,40,2,6);
    %
    %РњРµС‚РѕРґ 2
    %  L = superpixels(videoFrame,1000);
    %  points = bbox2points(bbox);
    %  roi = poly2mask(points(:,1),points(:,2),size(L,1),size(L,2));
    %  Face_cluster = grabcut(videoFrame,L,roi);
    %  maskedImage = videoFrame;
    %  maskedImage(repmat(~Face_cluster,[1 1 3])) = 0;
    %  imshow(maskedImage)
    %РњРµС‚РѕРґ 3
    % [BW,maskedImage] = segmentImage(videoFrame);

    [~, skmap] = skinmap(videoFrame);

    % Display the Skinmap data and draw the bounding box around the face.
%     figure(); imshow(skmap); title('Skinmap data');
%     rectangle('Position',bbox(1,:),'LineWidth',2,'EdgeColor',[1 1 0]);

    % Skin = bsxfun(@times, videoFrame, uint8(skmap));
    % Skin_gray = rgb2gray(Skin);
    % I_filtered = edge(Skin_gray, 'sobel');
    % figure; imshow(I_filtered);

    noseDetector = vision.CascadeObjectDetector('Nose', 'UseROI', true,'MergeThreshold',10);
    noseBBox     = step(noseDetector, videoFrame, bbox(1,:));
    %find center of nose Haar box
    % nx = noseBBox(1) + noseBBox(3)/2;
    % ny = noseBBox(2) + noseBBox(4)/2;
    nx = bbox(1,1);
    ny = bbox(1,2);

    %eyeBBox = step(EyeDetect, videoFrame, bbox(1,:));
%     rectangle('Position',noseBBox(1,:),'LineWidth',3,'EdgeColor',[1 1 0]);
%     rectangle('Position',noseBBox(1,:),'LineWidth',3,'EdgeColor',[1 1 0]);
%     hold on;
%     plot(nx,ny, 'Marker','+','Color','red','MarkerSize',10);

    % Create a tracker object.
    tracker = vision.HistogramBasedTracker;

    % Initialize the tracker histogram using the pixels from the nose.
    initializeObject(tracker, skmap, noseBBox(1,:));

    % Create a video player object for displaying video frames.
    %videoPlayer  = vision.VideoPlayer;

    frame = 0;
    numframes = int16(fix(vid_obj.FrameRate*vid_obj.Duration));
    % f = waitbar(0,'1','Name','Approximating pi...','CreateCancelBtn','setappdata(gcbf,''canceling'',1)');
    f = waitbar(numframes, "Начало обработки кадров");
    % Track the face over successive video frames until the video is finished.
    
    % Объявление матрицы RGB
    RGB = zeros(numframes, 3);
    file=fopen(strcat(VideoFile, '_RGB', '.txt'),'w'); 
    
    while  hasFrame(vid_obj)
        frame = frame + 1;
        waitbar(frame/numframes, f, sprintf("Обработка кадров [%d/%d] - %s", frame, numframes, VideoFile))

        % Extract the next video frame
        videoFrame = readFrame(vid_obj);
        
%             A_lin = rgb2lin(videoFrame);
%     percentiles = 10;
%     illuminant = illumgray(A_lin,percentiles);
%     B_lin = chromadapt(A_lin,illuminant,'ColorSpace','linear-rgb');
%     videoFrame = lin2rgb(B_lin);
         
%         if frame < 200
%              continue
%         end
%         if frame == 201
%             break
%         end
%         disp(frame);

        [~, skmap] = skinmap(videoFrame);

        % Track using the skinmap data
        bbox = step(tracker, skmap);

        % Apply the mask to the image
        Skin = bsxfun(@times, videoFrame, uint8(skmap));
        Face = imcrop(Skin, bbox);

        red = Face(:,:,1); 
        red = nonzeros(red);
        RGB(frame, 1) = sum(red)/length(red);
        green = Face(:,:,2);
        green = nonzeros(green);
        RGB(frame, 2) = sum(green)/length(green);%mean(green./length(green));
        blue = Face(:,:,3);
        blue = nonzeros(blue);
        RGB(frame, 3) = sum(blue)/length(blue);%mean(blue./length(blue));
       
%         disp(frame);
%         disp(RGB(frame, 1));
%         disp(RGB(frame, 2));
%         disp(RGB(frame, 3));
        
        if ((RGB(frame, 1) ~= 0) && (RGB(frame, 2) ~= 0) && (RGB(frame, 3) ~= 0))
            fprintf(file,'%f;%f;%f\n',RGB(frame, 1), RGB(frame, 2), RGB(frame, 3));
        end 

        % Insert a bounding box around the object being trackenamesd
        %videoOut = insertObjectAnnotation(Skin,'rectangle',bbox,'Face');
        % Display the annotated video frame using the video player object
        %step(videoPlayer, videoOut);
        %     disp(sprintf("[%d/%d]", frame, numframes));
    end
    delete(f);
    
fclose(file); 
%     writematrix([RGB(:, 1).',RGB(:, 2).',RGB(:, 3).'], strcat(VideoFile, '.txt'));
end 
% Release resources
%release(videoPlayer);