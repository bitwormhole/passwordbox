package com.bitwormhole.passwordbox.app;

import android.app.Activity;
import android.os.Bundle;

import androidx.annotation.Nullable;

import com.bitwormhole.passwordbox.app.ui.BaseActivity;
import com.bitwormhole.passwordbox.app.ui.DebugHomeActivity;
import com.bitwormhole.passwordbox.app.ui.LoginActivity;
import com.bitwormhole.passwordbox.app.ui.UnlockHomeActivity;
import com.bitwormhole.passwordbox.app.ui.tools.ActivityHelper;

public class MainActivity extends BaseActivity {

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.layout_main);

        ActivityHelper ah = new ActivityHelper(this);
        ah.setupButtonToOpenActivity(R.id.button_debug, DebugHomeActivity.class);
        ah.setupButtonToOpenActivity(R.id.button_login, LoginActivity.class);
        ah.setupButtonToOpenActivity(R.id.button_unlock, UnlockHomeActivity.class);

    }

}
