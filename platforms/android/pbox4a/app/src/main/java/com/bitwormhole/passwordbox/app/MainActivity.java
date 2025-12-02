package com.bitwormhole.passwordbox.app;

import android.app.Activity;
import android.os.Bundle;

import androidx.annotation.Nullable;

import com.bitwormhole.passwordbox.app.ui.BaseActivity;

public class MainActivity extends BaseActivity {

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.layout_main);
    }

}
