package com.bitwormhole.passwordbox.app.ui;

import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.TextView;

import androidx.annotation.Nullable;

import com.bitwormhole.passwordbox.app.R;
import com.bitwormhole.passwordbox.app.core.keybase.FingerPrint;
import com.bitwormhole.passwordbox.app.core.keybase.KeyPairAgent;
import com.bitwormhole.passwordbox.app.core.keybase.KeyPairHolder;
import com.bitwormhole.passwordbox.app.core.keybase.PublicKeys;
import com.bitwormhole.passwordbox.app.ui.tools.ActivityHelper;

import java.security.KeyPair;

public class DebugRootKeyActivity extends BaseActivity {

    private TextView mTextOutputSK;
    private TextView mTextOutputPK;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.layout_debug_root_key);

        mTextOutputPK = findViewById(R.id.text_output_public_key);
        mTextOutputSK = findViewById(R.id.text_output_secret_key);

        ActivityHelper ah = new ActivityHelper(this);
        ah.setupButtonToOpenActivity(R.id.button_load_public_key, this::handleClickButtonLoadPK);
        ah.setupButtonToOpenActivity(R.id.button_load_secret_key, this::handleClickButtonLoadSK);
    }


    private void handleClickButtonLoadPK(View v) {

        final String tag = this.getClass().getSimpleName();
        KeyPairHolder root = KeyPairAgent.getRoot();
        KeyPairHolder holder;

        if (root.exists()) {
            Log.i(tag, "load a existed key-pair");
            holder = root.load();
        } else {
            Log.i(tag, "generate a new key-pair");
            holder = root.generate();
        }

        KeyPair pair = holder.getPair();
        FingerPrint fp = PublicKeys.getFingerPrint(pair.getPublic());

        String msg = "public_key.finger-print = " + fp;
        Log.i(tag, msg);
    }

    private void handleClickButtonLoadSK(View v) {
        throw new RuntimeException("todo: no impl");
    }
}
