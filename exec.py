import os
import shutil
import sys
import threading

def multithread_wrapper(task):
    """Run target in another thread."""
    p = threading.Thread(target=task)
    p.start()

def wake_target_win():
    def task_block():
        cmd_check_batch_win = '''\
        cd quizzes/quiz{:0>2d} && \
        go build -o target.exe && \
        mv target.exe ../../_checkspace/target.exe && \
        cd ../../_checkspace/ && \
        target.exe
        '''
        os.system(cmd_check_batch_win.format(quiz_id))
    multithread_wrapper(task_block)


def wake_target():
    def task_block():
        cmd_check_batch = '''\
        cd quizzes/quiz{:0>2d} && \
        go build -o target && \
        mv target.exe ../../_checkspace/target && \
        cd ../../_checkspace/ && \
        ./target
        '''
        os.system(cmd_check_batch.format(quiz_id))
    multithread_wrapper(task_block)


def update_autograding():
    classroom_yml = action_template.format(quiz_id)
    os.makedirs('.github/workflows/', exist_ok=True)
    with open('.github/workflows/classroom.yml', 'w+') as f:
        f.write(classroom_yml)





def dispatch(arg):
    """dispatch

    Attention:
        All branches started with function call
        `init_quiz_id()`, need the extra argument : [id]
    """
    if arg == 'init':
        init_quiz_id()
        update_autograding()

    elif arg == '__check':  # called by classroom action : education/autograding@v1
        init_quiz_id()
        check()

    elif arg == 'test':  # used for local test
        init_quiz_id()
        check_locally()

    elif arg == 'help':
        print(help_msg)

    elif arg == 'clean':
        clean()

    else:
        print('[ERROR] Unknown arg.')


if __name__ == '__main__':
    if len(sys.argv) < 2:
        print('[ERROR] Args were not enough.')
        exit(1)

    dispatch(sys.argv[1])
    exit(0)
