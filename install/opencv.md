## install dep
```
apt install cmake make gcc g++ pkg-config
```

## install opencv
```
cmake -D CMAKE_BUILD_TYPE=RELEASE \
      -D WITH_IPP=OFF \
      -D WITH_OPENGL=OFF \
      -D WITH_QT=OFF \
      -D BUILD_EXAMPLES=OFF \
      -D BUILD_TESTS=OFF \
      -D BUILD_PERF_TESTS=OFF  \
      -D BUILD_opencv_java=OFF \
      -D BUILD_opencv_python=OFF \
      -D BUILD_opencv_python2=OFF \
      -D BUILD_opencv_python3=OFF \
      -D OPENCV_GENERATE_PKGCONFIG=ON \
      -D OPENCV_ENABLE_NONFREE=ON \
      -D OPENCV_EXTRA_MODULES_PATH=/usr/local/opencv_contrib/modules \
      -D CMAKE_INSTALL_PREFIX=/usr/local/opencv ..
make
make install
ldconfig
```

## set env
```
echo "export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/opencv/lib/pkgconfig" >> ~/.bashrc
echo "export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/opencv/lib" >> ~/.bashrc
source ~/.bashrc
```

## make java env (jar & so)
```
cmake -D CMAKE_BUILD_TYPE=RELEASE \
      -D BUILD_SHARED_LIBS=OFF \
      -D BUILD_FAT_JAVA_LIB=ON \
      -D CMAKE_INSTALL_PREFIX=/usr/local/opencv ..
make
make install
```