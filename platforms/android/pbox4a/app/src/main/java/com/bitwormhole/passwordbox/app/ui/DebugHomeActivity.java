package com.bitwormhole.passwordbox.app.ui;

import android.os.Bundle;

import androidx.annotation.Nullable;

import com.bitwormhole.passwordbox.app.R;

public class DebugHomeActivity extends BaseActivity {

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.layout_debug_home);
    }

}
