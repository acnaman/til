import { Component, Element, Prop, State, Event, EventEmitter } from "@stencil/core";

@Component({
    tag: "app-textarea",
    styleUrl: "app-textarea.scss"
})

export class AppTextarea {
    @Prop() placeholder: string = "";
    @Prop() btText: string = "";

    @Prop({ connect: "ion-loading-controller" })
    loadingCtrl: HTMLIonLoadingControllerElement;

    @Element() el: HTMLElement;

    @State() text: string = "";
    @State() validate: boolean = true;

    @Event() sendBtClicked: EventEmitter;

    textInput(el) {
        this.text = el.srcElement.value;
        if (this.text.length <= 0 || 100 < this.text.length){
            this.validate = true;
        } else {
            this.validate = false;
        }
    }

    async send(){
        this.sendBtClicked.emit(this.text);
    }
    
    render () {
        return (
            <div class="textarea-wrapper">
                <textarea
                    id="textarea"
                    placeholder={this.placeholder}
                    onInput={e => this.textInput(e)}
                />
                <div class="number-of-characters">{this.text.length}/100</div>
                <ion-button class="send-button" disabled={this.validate}>
                    <ion-icon slot="start" name="send"/>
                    {this.btText}
                </ion-button>
            </div>
        );
    }
}


